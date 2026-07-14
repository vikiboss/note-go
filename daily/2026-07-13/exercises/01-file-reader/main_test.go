package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadText(t *testing.T) {
	path := filepath.Join(t.TempDir(), "sample.txt")
	if err := os.WriteFile(path, []byte("hello Go"), 0o600); err != nil {
		t.Fatal(err)
	}

	got, err := readText(path)
	if err != nil {
		t.Fatal(err)
	}
	if got != "hello Go" {
		t.Fatalf("readText() = %q, want %q", got, "hello Go")
	}
}
