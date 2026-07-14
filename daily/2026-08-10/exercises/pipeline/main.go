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

type indexedJob struct {
	index int
	input int
}

type indexedResult struct {
	index  int
	result Result
}

func Squares(ctx context.Context, inputs []int, workers int) ([]Result, error) {
	if workers < 1 {
		return nil, fmt.Errorf("workers must be positive")
	}
	jobs := make(chan indexedJob)
	results := make(chan indexedResult)
	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for job := range jobs {
				select {
				case results <- indexedResult{
					index:  job.index,
					result: Result{Input: job.input, Value: job.input * job.input},
				}:
				case <-ctx.Done():
					return
				}
			}
		}()
	}
	go func() {
		defer close(jobs)
		for index, input := range inputs {
			select {
			case jobs <- indexedJob{index: index, input: input}:
			case <-ctx.Done():
				return
			}
		}
	}()
	go func() {
		wg.Wait()
		close(results)
	}()

	out := make([]Result, len(inputs))
	for result := range results {
		out[result.index] = result.result
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
