package counter

import "sync"

type Counter struct {
	mu    sync.RWMutex
	value int64
}

func (c *Counter) Add(delta int64) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += delta
	return c.value
}

func (c *Counter) Value() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.value
}
