package main

import (
	"os"
	"path/filepath"
	"strings"
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

func TestReadTextWrapsPathOnFailure(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.txt")
	_, err := readText(path)
	if err == nil || !strings.Contains(err.Error(), path) {
		t.Fatalf("error = %v, want path %q", err, path)
	}
}

func TestSummarizeTextDistinguishesBytesAndRunes(t *testing.T) {
	got := summarizeText("Go 你好\n")
	want := textStats{Bytes: 10, Runes: 6, Words: 2, Lines: 1}
	if got != want {
		t.Fatalf("summarizeText() = %#v, want %#v", got, want)
	}
}
