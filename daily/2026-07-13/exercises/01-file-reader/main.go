package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func readText(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %q: %w", path, err)
	}
	return string(data), nil
}

func main() {
	path := flag.String("file", "", "text file to read")
	flag.Parse()
	if *path == "" {
		fmt.Fprintln(os.Stderr, "usage: go run ./exercises/01-file-reader -file path/to/file")
		os.Exit(2)
	}

	text, err := readText(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	words := len(strings.Fields(text))
	fmt.Printf("bytes: %d, words: %d\n%s", len(text), words, text)
}
