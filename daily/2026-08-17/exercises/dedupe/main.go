package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"sort"
	"sync"
)

type hashResult struct {
	path string
	sum  string
	err  error
}

func FindDuplicates(ctx context.Context, paths []string, workers int) (map[string][]string, error) {
	if workers < 1 {
		return nil, fmt.Errorf("workers must be positive")
	}
	jobs := make(chan string)
	results := make(chan hashResult)
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range jobs {
				data, err := os.ReadFile(path)
				r := hashResult{path: path, err: err}
				if err == nil {
					r.sum = fmt.Sprintf("%x", sha256.Sum256(data))
				}
				select {
				case results <- r:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	go func() {
		defer close(jobs)
		for _, path := range paths {
			select {
			case jobs <- path:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	groups := make(map[string][]string)
	var firstErr error
	for result := range results {
		if result.err != nil {
			if firstErr == nil {
				firstErr = fmt.Errorf("hash %q: %w", result.path, result.err)
			}
			continue
		}
		groups[result.sum] = append(groups[result.sum], result.path)
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if firstErr != nil {
		return nil, firstErr
	}
	for sum, group := range groups {
		if len(group) < 2 {
			delete(groups, sum)
			continue
		}
		sort.Strings(group)
	}
	return groups, nil
}

func main() {
	groups, err := FindDuplicates(context.Background(), os.Args[1:], 4)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for sum, paths := range groups {
		fmt.Printf("%.12s %v\n", sum, paths)
	}
}
