package main

import (
	"context"
	"errors"
	"sort"
	"testing"
)

func TestSquares(t *testing.T) {
	got, err := Squares(context.Background(), []int{3, 1, 2}, 2)
	if err != nil {
		t.Fatal(err)
	}
	sort.Slice(got, func(i, j int) bool { return got[i].Input < got[j].Input })
	want := []int{1, 4, 9}
	for i, result := range got {
		if result.Value != want[i] {
			t.Fatalf("result[%d] = %#v", i, result)
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
