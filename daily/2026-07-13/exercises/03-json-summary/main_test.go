package main

import "testing"

func TestSummarize(t *testing.T) {
	total, done, err := summarize([]byte(`[{"title":"learn Go","done":true},{"title":"write tests","done":false}]`))
	if err != nil {
		t.Fatal(err)
	}
	if total != 2 || done != 1 {
		t.Fatalf("summarize() = (%d, %d), want (2, 1)", total, done)
	}
}
