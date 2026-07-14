package main

import (
	"context"
	"errors"
	"testing"
)

func TestSquares(t *testing.T) {
	got, err := Squares(context.Background(), []int{3, 1, 2}, 2)
	if err != nil {
		t.Fatal(err)
	}
	want := []Result{{Input: 3, Value: 9}, {Input: 1, Value: 1}, {Input: 2, Value: 4}}
	for i, result := range got {
		if result != want[i] {
			t.Fatalf("result[%d] = %#v, want %#v", i, result, want[i])
		}
	}
}

func TestSquaresCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := Squares(ctx, []int{1, 2, 3}, 2)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
}

func TestSquaresRejectsInvalidWorkerCount(t *testing.T) {
	if _, err := Squares(context.Background(), []int{1}, 0); err == nil {
		t.Fatal("expected worker validation error")
	}
}
