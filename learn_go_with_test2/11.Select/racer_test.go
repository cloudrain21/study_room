package racer

import (
    "fmt"
    "net/http"
    "testing"
    "time"
    "net/http/httptest"
)

func TestRacer(t *testing.T) {
    // t.Run("real web : this test may fail", func(t *testing.T) {
    //     fastUrl := "http://www.facebook.com"
    //     slowUrl := "www.quii.co.uk"

    //     got := Racer(slowUrl, fastUrl)
    //     want := fastUrl

    //     if got != want {
    //         t.Errorf("got %q, want %q", got, want)
    //     }
    // })

    t.Run("mock web", func(t *testing.T) {
        slowServer := makeDelayServer(20 * time.Millisecond)
        fastServer := makeDelayServer(0 * time.Millisecond)

        defer slowServer.Close()
        defer fastServer.Close()

        slowUrl := slowServer.URL
        fastUrl := fastServer.URL

        want := fastUrl
        got := Racer(slowUrl, fastUrl)

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })

    t.Run("test web parallel", func(t *testing.T) {
        slowServer := makeDelayServer(20 * time.Millisecond)
        fastServer := makeDelayServer(0 * time.Millisecond)

        defer slowServer.Close()
        defer fastServer.Close()

        slowUrl := slowServer.URL
        fastUrl := fastServer.URL

        want := fastUrl
        got := RacerParallel(slowUrl, fastUrl)

        if got != want {
            t.Errorf("got %q, want %q", got, want)
        }
    })

    t.Run("timeout", func(t *testing.T) {
        serverA := makeDelayServer(2 * time.Second)
        serverB := makeDelayServer(3 * time.Second)

        defer serverA.Close()
        defer serverB.Close()

        TimeoutSec = 1 * time.Second
        _, err := RacerTimeout(serverA.URL, serverB.URL)
        if err == nil {
            t.Error("expected error but didn't get one")
        }
    })
}

func ExampleRacer() {
    serverA := makeDelayServer(2 * time.Second)
    serverB := makeDelayServer(3 * time.Second)

    defer serverA.Close()
    defer serverB.Close()

    TimeoutSec = 1 * time.Second
    _, err := RacerTimeout(serverA.URL, serverB.URL)
    if err != nil {
        fmt.Println(err)
    }

    //Output:
    //timed out waiting for http and http
}

func makeDelayServer(d time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(d)
        w.WriteHeader(http.StatusOK)
    }))
}
