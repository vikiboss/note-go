package taskstore

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

type Task struct {
	ID             int    `json:"id"`
	IdempotencyKey string `json:"idempotency_key"`
	Payload        string `json:"payload"`
	Status         string `json:"status"`
}

type Store struct {
	mu    sync.Mutex
	path  string
	tasks []Task
}

func Open(path string) (*Store, error) {
	s := &Store{path: path}
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return s, nil
		}
		return nil, fmt.Errorf("open store: %w", err)
	}
	if err := json.Unmarshal(data, &s.tasks); err != nil {
		return nil, fmt.Errorf("decode store: %w", err)
	}
	return s, nil
}

func (s *Store) Create(key, payload string) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if key == "" {
		return Task{}, fmt.Errorf("idempotency key is required")
	}
	for _, task := range s.tasks {
		if task.IdempotencyKey == key {
			return task, nil
		}
	}
	task := Task{ID: len(s.tasks) + 1, IdempotencyKey: key, Payload: payload, Status: "pending"}
	next := append(append([]Task(nil), s.tasks...), task)
	if err := s.persist(next); err != nil {
		return Task{}, err
	}
	s.tasks = next
	return task, nil
}

func (s *Store) List() []Task {
	s.mu.Lock()
	defer s.mu.Unlock()
	return append([]Task(nil), s.tasks...)
}

func (s *Store) persist(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return fmt.Errorf("encode store: %w", err)
	}
	if err := os.MkdirAll(filepath.Dir(s.path), 0o755); err != nil {
		return fmt.Errorf("create store directory: %w", err)
	}
	tmp, err := os.CreateTemp(filepath.Dir(s.path), ".tasks-*")
	if err != nil {
		return fmt.Errorf("create temporary store: %w", err)
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		return fmt.Errorf("write temporary store: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return fmt.Errorf("sync temporary store: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("close temporary store: %w", err)
	}
	if err := os.Rename(tmpName, s.path); err != nil {
		return fmt.Errorf("replace store: %w", err)
	}
	return nil
}
