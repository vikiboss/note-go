package main

import (
	"strings"
	"testing"
)

func TestSummarize(t *testing.T) {
	total, done, err := summarize([]byte(`[{"title":"learn Go","done":true},{"title":"write tests","done":false}]`))
	if err != nil {
		t.Fatal(err)
	}
	if total != 2 || done != 1 {
		t.Fatalf("summarize() = (%d, %d), want (2, 1)", total, done)
	}
}

func TestSummarizeRejectsInvalidInput(t *testing.T) {
	tests := []struct {
		name string
		data string
		want string
	}{
		{name: "malformed JSON", data: "[", want: "decode tasks"},
		{name: "blank title", data: `[{"title":"  ","done":false}]`, want: "title is required"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := summarize([]byte(tt.data))
			if err == nil || !strings.Contains(err.Error(), tt.want) {
				t.Fatalf("error = %v, want substring %q", err, tt.want)
			}
		})
	}
}
