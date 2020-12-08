package racer

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "time"
)

func TestRacer(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        slowServer := makeDelayServer(600 * time.Millisecond)
        fastServer := makeDelayServer(500 * time.Millisecond)

        defer slowServer.Close()
        defer fastServer.Close()

        slowURL := slowServer.URL
        fastURL := fastServer.URL

        //want := fastURL
        //got := Racer(slowURL, fastURL)
        _, err := RacerConcurrency(slowURL, fastURL, 100 * time.Millisecond)

        if err == nil {
            t.Error("expected error but...")
        }

        slowServer.Close()
        fastServer.Close()
    })
}

func makeDelayServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}
