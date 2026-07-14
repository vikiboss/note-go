package taskstore

import (
	"errors"
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

func TestTaskStatusTransitionsArePersisted(t *testing.T) {
	path := filepath.Join(t.TempDir(), "tasks.json")
	store, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	task, err := store.Create("request-1", "work")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := store.UpdateStatus(task.ID, "done"); err == nil {
		t.Fatal("pending -> done should be rejected")
	}
	if _, err := store.UpdateStatus(task.ID, "running"); err != nil {
		t.Fatal(err)
	}
	if _, err := store.UpdateStatus(task.ID, "done"); err != nil {
		t.Fatal(err)
	}
	if _, err := store.UpdateStatus(task.ID, "running"); err == nil {
		t.Fatal("terminal task should not restart")
	}
	reopened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if got := reopened.List()[0].Status; got != "done" {
		t.Fatalf("status after reopen = %q, want done", got)
	}
	if _, err := store.UpdateStatus(999, "running"); !errors.Is(err, ErrTaskNotFound) {
		t.Fatalf("error = %v, want ErrTaskNotFound", err)
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
