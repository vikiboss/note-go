package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanIsStableAndIgnoresGit(t *testing.T) {
	root := t.TempDir()
	for name, data := range map[string]string{"b.txt": "b", "nested/a.txt": "a", ".git/secret": "x"} {
		path := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	got, err := Scan(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got[0].Path != "b.txt" || got[1].Path != "nested/a.txt" {
		t.Fatalf("entries = %#v", got)
	}
	if got[0].Hash == "" || got[1].Hash == "" {
		t.Fatal("missing content hash")
	}
}
