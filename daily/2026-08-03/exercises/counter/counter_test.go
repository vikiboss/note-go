package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	tests := []struct {
		name   string
		deltas []int64
		want   int64
	}{
		{name: "zero value works"},
		{name: "positive", deltas: []int64{1, 2, 3}, want: 6},
		{name: "mixed", deltas: []int64{5, -2}, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var c Counter
			for _, delta := range tt.deltas {
				c.Add(delta)
			}
			if got := c.Value(); got != tt.want {
				t.Fatalf("Value() = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestCounterConcurrent(t *testing.T) {
	var c Counter
	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 100 {
				c.Add(1)
			}
		}()
	}
	wg.Wait()
	if got := c.Value(); got != 10_000 {
		t.Fatalf("Value() = %d, want 10000", got)
	}
}

func BenchmarkCounterAdd(b *testing.B) {
	var c Counter
	b.ResetTimer()
	for range b.N {
		c.Add(1)
	}
}
