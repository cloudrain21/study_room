package server

import (
    "fmt"
    "net/http"
)

type Store interface {
    Fetch() string
    Cancel()
}

type StubStore struct {
    response string
    cancelled bool
}

func (s *StubStore)Fetch() string {
    time.Sleep(100 * time.Millisecond)
    return s.response
}

func (s *StubStore)Cancel() {
    s.cancelled = true
}

func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, store.Fetch())
    }
}
