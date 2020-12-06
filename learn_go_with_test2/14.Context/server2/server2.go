package server2

import (
    "context"
    "fmt"
    "net/http"
)

type Store interface {
    Fetch(ctx context.Context) (string,error)
}

func Server(store Store) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        s, err := store.Fetch(r.Context())

        if err != nil {
            // response not witten
            return
        }

        fmt.Fprintf(w, s)
    }
}
