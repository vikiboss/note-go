package filesync

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

type cancelReader struct {
	cancel context.CancelFunc
	done   bool
}

func (r *cancelReader) Read(p []byte) (int, error) {
	if r.done {
		return 0, errors.New("unexpected second read")
	}
	r.done = true
	copy(p, "chunk")
	r.cancel()
	return len("chunk"), nil
}

func TestSyncDryRunAndApply(t *testing.T) {
	source, target := t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(source, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}
	sourcePath := filepath.Join(source, "nested", "a.txt")
	if err := os.WriteFile(sourcePath, []byte("hello"), 0o600); err != nil {
		t.Fatal(err)
	}
	actions, err := Sync(context.Background(), source, target, true)
	if err != nil || len(actions) != 1 || actions[0].Kind != "copy" {
		t.Fatalf("dry run = %#v, %v", actions, err)
	}
	targetPath := filepath.Join(target, "nested", "a.txt")
	if _, err := os.Stat(targetPath); !os.IsNotExist(err) {
		t.Fatalf("dry run changed target: %v", err)
	}
	actions, err = Sync(context.Background(), source, target, false)
	if err != nil || len(actions) != 1 {
		t.Fatalf("apply = %#v, %v", actions, err)
	}
	data, err := os.ReadFile(targetPath)
	if err != nil || string(data) != "hello" {
		t.Fatalf("target = %q, %v", data, err)
	}
	actions, err = Sync(context.Background(), source, target, false)
	if err != nil || len(actions) != 0 {
		t.Fatalf("second sync = %#v, %v", actions, err)
	}
}

func TestCopyWithContextCanCancelLargeFileMidStream(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var dst bytes.Buffer
	_, err := copyWithContext(ctx, &dst, &cancelReader{cancel: cancel})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
	if dst.String() != "chunk" {
		t.Fatalf("partial output = %q", dst.String())
	}
}
