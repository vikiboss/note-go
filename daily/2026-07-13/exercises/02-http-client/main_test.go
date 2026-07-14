package main

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestFetchRejectsNon2xxAndKeepsStatus(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: http.StatusTooManyRequests,
			Body:       io.NopCloser(strings.NewReader("slow down")),
			Header:     make(http.Header),
		}, nil
	})}
	status, body, err := fetchWithClient(client, "https://example.test/limited")
	if status != http.StatusTooManyRequests || body != nil {
		t.Fatalf("result = (%d, %q), want (429, nil)", status, body)
	}
	var statusErr *HTTPStatusError
	if !errors.As(err, &statusErr) || statusErr.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("error = %#v, want HTTPStatusError(429)", err)
	}
}

func TestFetchWrapsTransportError(t *testing.T) {
	sentinel := errors.New("network down")
	client := &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		return nil, sentinel
	})}
	_, _, err := fetchWithClient(client, "https://example.test")
	if !errors.Is(err, sentinel) {
		t.Fatalf("error = %v, want wrapped sentinel", err)
	}
}

func TestFetch(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		if r.URL.Path != "/health" {
			t.Errorf("path = %q, want %q", r.URL.Path, "/health")
		}
		return &http.Response{
			StatusCode: http.StatusCreated,
			Body:       io.NopCloser(strings.NewReader("created")),
			Header:     make(http.Header),
		}, nil
	})}

	status, body, err := fetchWithClient(client, "https://example.test/health")
	if err != nil {
		t.Fatal(err)
	}
	if status != http.StatusCreated {
		t.Fatalf("status = %d, want %d", status, http.StatusCreated)
	}
	if string(body) != "created" {
		t.Fatalf("body = %q, want %q", body, "created")
	}
}
