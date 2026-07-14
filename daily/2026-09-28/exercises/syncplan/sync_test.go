package syncplan

import (
	"bytes"
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

type cancelAfterFirstRead struct {
	cancel context.CancelFunc
	done   bool
}

func (r *cancelAfterFirstRead) Read(p []byte) (int, error) {
	if r.done {
		return 0, errors.New("reader should not be called again")
	}
	r.done = true
	copy(p, "partial")
	r.cancel()
	return len("partial"), nil
}

func TestPlan(t *testing.T) {
	source := []Entry{{Path: "b", Hash: "new"}, {Path: "a", Hash: "same"}, {Path: "c", Hash: "x"}}
	target := []Entry{{Path: "a", Hash: "same"}, {Path: "b", Hash: "old"}, {Path: "extra", Hash: "keep"}}
	got := Plan(source, target)
	want := []Action{{Path: "b", Kind: "update"}, {Path: "c", Kind: "copy"}}
	if len(got) != len(want) {
		t.Fatalf("Plan() = %#v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("action[%d] = %#v, want %#v", i, got[i], want[i])
		}
	}
}

func TestCopyWithContextStopsDuringFile(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	var dst bytes.Buffer
	_, err := copyWithContext(ctx, &dst, &cancelAfterFirstRead{cancel: cancel})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
	if dst.String() != "partial" {
		t.Fatalf("partial output = %q", dst.String())
	}
}

func TestApplyCopiesAndHonorsCancellation(t *testing.T) {
	source, target := t.TempDir(), t.TempDir()
	if err := os.MkdirAll(filepath.Join(source, "nested"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(source, "nested", "a"), []byte("new"), 0o600); err != nil {
		t.Fatal(err)
	}
	actions := []Action{{Path: "nested/a", Kind: "copy"}}
	if err := Apply(context.Background(), source, target, actions); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(target, "nested", "a"))
	if err != nil || string(data) != "new" {
		t.Fatalf("target = %q, %v", data, err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := Apply(ctx, source, target, actions); !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
}
