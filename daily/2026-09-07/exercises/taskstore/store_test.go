package taskstore

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCreateIsIdempotentAndPersistent(t *testing.T) {
	path := filepath.Join(t.TempDir(), "data", "tasks.json")
	store, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	first, err := store.Create("request-1", "send report")
	if err != nil {
		t.Fatal(err)
	}
	second, err := store.Create("request-1", "different payload is ignored")
	if err != nil {
		t.Fatal(err)
	}
	if first != second || len(store.List()) != 1 {
		t.Fatalf("tasks are not idempotent: %#v %#v", first, second)
	}
	reopened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := reopened.List(); len(got) != 1 || got[0] != first {
		t.Fatalf("reopened tasks = %#v", got)
	}
}

func TestOpenRejectsCorruptData(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	if err := os.WriteFile(path, []byte("{"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := Open(path); err == nil {
		t.Fatal("expected corrupt data error")
	}
}
