package racer

import (
    "fmt"
    "net/http"
    "strings"
    "time"
)

///
func Racer(a, b string) string {
    if measureResponseTime(a) > measureResponseTime(b) {
        return b
    }
    return a
}

func measureResponseTime(url string) time.Duration {
    startA := time.Now()
    http.Get(url)
    return time.Since(startA)
}

///
func RacerParallel(a, b string) string {
    select {
        case <- ping(a):
            return a
        case <- ping(b):
            return b
    }
}

func ping(url string) chan struct{} {
    ch := make(chan struct{})

    go func() {
        http.Get(url)
        close(ch)
    }()

    return ch
}

///
var TimeoutSec time.Duration

func RacerTimeout(a, b string) (string, error) {
    return ConfigurableRacer(a, b, TimeoutSec)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
    select {
        case <- ping(a):
            return a, nil
        case <- ping(b):
            return b, nil
        case <- time.After(timeout):
            return "", fmt.Errorf("timed out waiting for %s and %s", strings.Split(a,":")[0], strings.Split(b,":")[0])
    }
}
