package analyzer

import (
	"strings"
	"unicode"
)

type Stats struct {
	Words int
	Lines int
	Bytes int
}

func Analyze(text string) Stats {
	stats := Stats{Bytes: len(text)}
	if text != "" {
		stats.Lines = strings.Count(text, "\n") + 1
	}
	inWord := false
	for _, r := range text {
		if unicode.IsSpace(r) {
			inWord = false
			continue
		}
		if !inWord {
			stats.Words++
			inWord = true
		}
	}
	return stats
}

func AnalyzeWithFields(text string) Stats {
	stats := Stats{Bytes: len(text), Words: len(strings.Fields(text))}
	if text != "" {
		stats.Lines = strings.Count(text, "\n") + 1
	}
	return stats
}
