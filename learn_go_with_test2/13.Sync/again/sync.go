package sync

import (
    "sync"
)

type Counter struct {
    sync.Mutex
    val int
}

func NewCounter() *Counter {
    return &Counter{}
}

func (c *Counter)Inc() {
    c.Lock()
    defer c.Unlock()

    c.val++
}

func (c *Counter)Value() int {
    return c.val
}

