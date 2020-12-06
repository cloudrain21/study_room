package server

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
        cnxt := r.Context()

        data := make(chan string, 1)

        go func() {
            data <- store.Fetch()
            close(data)
        }()

        select {
            case d := <- data:
                fmt.Fprintf(w, d)
            case <- cnxt.Done():
                store.Cancel()
        }
    }
}
