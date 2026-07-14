package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type textStats struct {
	Bytes int
	Runes int
	Words int
	Lines int
}

func readText(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %q: %w", path, err)
	}
	return string(data), nil
}

func summarizeText(text string) textStats {
	stats := textStats{
		Bytes: len(text),
		Runes: utf8.RuneCountInString(text),
		Words: len(strings.Fields(text)),
	}
	if text != "" {
		stats.Lines = strings.Count(text, "\n") + 1
		if strings.HasSuffix(text, "\n") {
			stats.Lines--
		}
	}
	return stats
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

	stats := summarizeText(text)
	fmt.Printf("bytes: %d, runes: %d, words: %d, lines: %d\n%s", stats.Bytes, stats.Runes, stats.Words, stats.Lines, text)
}
