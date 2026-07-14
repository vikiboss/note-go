package analyzer

import (
	"strings"
	"testing"
)

func TestAnalyze(t *testing.T) {
	tests := []struct {
		text string
		want Stats
	}{
		{text: "", want: Stats{}},
		{text: "hello Go", want: Stats{Words: 2, Lines: 1, Bytes: 8}},
		{text: "你好 Go\n并发", want: Stats{Words: 3, Lines: 2, Bytes: 16}},
	}
	for _, tt := range tests {
		if got := Analyze(tt.text); got != tt.want {
			t.Fatalf("Analyze(%q) = %#v, want %#v", tt.text, got, tt.want)
		}
	}
}

var benchmarkText = strings.Repeat("Go makes concurrency explicit.\n", 1000)

func BenchmarkAnalyze(b *testing.B) {
	for range b.N {
		_ = Analyze(benchmarkText)
	}
}

func BenchmarkAnalyzeWithFields(b *testing.B) {
	for range b.N {
		_ = AnalyzeWithFields(benchmarkText)
	}
}
