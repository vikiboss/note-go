package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

const maxJSONLLine = 1 << 20

type Event struct {
	Level string `json:"level"`
}

type Summary struct {
	Levels map[string]int `json:"levels"`
	Total  int            `json:"total"`
}

func Run(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 64<<10), maxJSONLLine)
	summary := Summary{Levels: make(map[string]int)}
	line := 0
	for scanner.Scan() {
		line++
		var event Event
		if err := json.Unmarshal(scanner.Bytes(), &event); err != nil {
			return fmt.Errorf("line %d: decode JSON: %w", line, err)
		}
		if event.Level == "" {
			return fmt.Errorf("line %d: level is required", line)
		}
		summary.Levels[event.Level]++
		summary.Total++
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("scan input: %w", err)
	}
	// encoding/json emits string map keys in lexical order, so CLI output is stable.
	return json.NewEncoder(w).Encode(summary)
}

func main() {
	if err := Run(os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
