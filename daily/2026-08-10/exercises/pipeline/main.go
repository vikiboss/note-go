package main

import (
	"context"
	"fmt"
	"sync"
)

type Result struct {
	Input int
	Value int
}

func Squares(ctx context.Context, inputs []int, workers int) ([]Result, error) {
	if workers < 1 {
		return nil, fmt.Errorf("workers must be positive")
	}
	jobs := make(chan int)
	results := make(chan Result)
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				select {
				case results <- Result{Input: n, Value: n * n}:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	go func() {
		defer close(jobs)
		for _, n := range inputs {
			select {
			case jobs <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	var out []Result
	for result := range results {
		out = append(out, result)
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return out, nil
}

func main() {
	results, _ := Squares(context.Background(), []int{1, 2, 3, 4}, 2)
	fmt.Println(results)
}
