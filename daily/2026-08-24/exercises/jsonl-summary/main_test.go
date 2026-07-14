package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	in := strings.NewReader("{\"level\":\"info\"}\n{\"level\":\"error\"}\n{\"level\":\"info\"}\n")
	var out bytes.Buffer
	if err := Run(in, &out); err != nil {
		t.Fatal(err)
	}
	if got, want := out.String(), "{\"levels\":{\"error\":1,\"info\":2},\"total\":3}\n"; got != want {
		t.Fatalf("output = %q, want %q", got, want)
	}
}

func TestRunIncludesLineNumber(t *testing.T) {
	var out bytes.Buffer
	err := Run(strings.NewReader("{\"level\":\"info\"}\nnot-json\n"), &out)
	if err == nil || !strings.Contains(err.Error(), "line 2") {
		t.Fatalf("error = %v, want line 2", err)
	}
}
