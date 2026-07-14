package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string) (int, []byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	return fetchWithClient(client, url)
}

func fetchWithClient(client *http.Client, url string) (int, []byte, error) {
	resp, err := client.Get(url)
	if err != nil {
		return 0, nil, fmt.Errorf("GET %s: %w", url, err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 4<<10))
	if err != nil {
		return resp.StatusCode, nil, fmt.Errorf("read response: %w", err)
	}
	return resp.StatusCode, body, nil
}

func main() {
	url := flag.String("url", "https://example.com", "URL to fetch")
	flag.Parse()

	status, body, err := fetch(*url)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("status: %d\nbody (first 4 KiB):\n%s\n", status, body)
}
