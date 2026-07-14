package filesync

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

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
