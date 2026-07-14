package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

type task struct {
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

func summarize(data []byte) (total, done int, err error) {
	var tasks []task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return 0, 0, fmt.Errorf("decode tasks: %w", err)
	}
	for _, t := range tasks {
		if strings.TrimSpace(t.Title) == "" {
			return 0, 0, fmt.Errorf("decode tasks: title is required")
		}
		if t.Done {
			done++
		}
	}
	return len(tasks), done, nil
}

func main() {
	path := flag.String("file", "", "JSON task list")
	flag.Parse()
	if *path == "" {
		fmt.Fprintln(os.Stderr, "usage: go run ./exercises/03-json-summary -file tasks.json")
		os.Exit(2)
	}
	data, err := os.ReadFile(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	total, done, err := summarize(data)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("tasks: %d, complete: %d, remaining: %d\n", total, done, total-done)
}
