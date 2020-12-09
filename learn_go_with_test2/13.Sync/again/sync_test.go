package sync

import (
    "sync"
    "testing"
)

func TestCounter(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        counter := &Counter{}

        counter.Inc()
        counter.Inc()
        counter.Inc()

        assertCounter(t, counter, 3)
    })

    t.Run("wait group counter", func(t *testing.T) {
        counter := NewCounter()
        want := 1000

        var wg sync.WaitGroup

        for i:=0; i<want; i++ {
            wg.Add(1)

            go func(wg *sync.WaitGroup) {
                defer wg.Done()
                counter.Inc()
            }(&wg)
        }

        wg.Wait()

        assertCounter(t, counter, want)
    })
}

func assertCounter(t *testing.T, counter *Counter, want int) {
    t.Helper()

    if counter.Value() != want {
        t.Errorf("got %d want %d", counter.Value(), want)
    }
}
