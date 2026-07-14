package main

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestFindDuplicates(t *testing.T) {
	dir := t.TempDir()
	var paths []string
	for name, data := range map[string]string{"a": "same", "b": "same", "c": "other"} {
		path := filepath.Join(dir, name)
		if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
			t.Fatal(err)
		}
		paths = append(paths, path)
	}
	groups, err := FindDuplicates(context.Background(), paths, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(groups) != 1 {
		t.Fatalf("groups = %#v, want one duplicate group", groups)
	}
	for _, group := range groups {
		if len(group) != 2 {
			t.Fatalf("group = %#v, want two paths", group)
		}
	}
}

func TestFindDuplicatesReportsPath(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "missing")
	_, err := FindDuplicates(context.Background(), []string{missing}, 1)
	if err == nil {
		t.Fatal("expected error")
	}
}
