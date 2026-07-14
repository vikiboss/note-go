package syncplan

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"
)

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
