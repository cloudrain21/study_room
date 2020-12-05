package counter

import (
    "sync"
    "testing"
)

func TestCounter(t *testing.T) {
    t.Run("count 3", func(t *testing.T) {
        counter := Counter{}

        counter.Inc()
        counter.Inc()
        counter.Inc()

        assertCounter(t, counter.Value(), 3)
    })

    t.Run("concurrent", func(t *testing.T) {
        counter := NewCounter()
        var wg sync.WaitGroup

        want := 1000

        for i:=0; i<want; i++ {
            wg.Add(1)
            go func(wg *sync.WaitGroup) {
                counter.Inc()
                wg.Done()
            }(&wg)
        }
        wg.Wait()

        //got := counter.Value()
        //assertCounter(t, got, want)

        assertCounter2(t, counter, want)
    })
}

func NewCounter() *Counter {
    return &Counter{}
}

func assertCounter(t *testing.T, got, want int) {
    t.Helper()

    if got != want {
        t.Errorf("got %v want %v", got, want)
    }
}

// bad idea to copy lock value in Counter - use pointer of counter
func assertCounter2(t *testing.T, counter *Counter, want int) {
    t.Helper()

    if counter.Value() != want {
        t.Errorf("got %v want %v", counter.Value(), want)
    }
}
