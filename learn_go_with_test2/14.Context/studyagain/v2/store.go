package store

import (
    "fmt"
    "net/http"
)

type Store interface {
    Fetch() string
    Cancel()
}

func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()

        dataChan := make(chan string, 1)

        go func() {
            dataChan <- store.Fetch()
        }()

        select {
        case <- ctx.Done():
            store.Cancel()
        case data := <- dataChan:
            fmt.Fprintf(w, data)
        }
    }
}
