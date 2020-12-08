package racer

import (
    "fmt"
    "net/http"
    "time"
)

func Racer(a, b string) (string, error) {
    aDuration := measureTime(a)
    bDuration := measureTime(b)

    if aDuration < bDuration {
        return a, nil
    }
    return b, nil
}

func RacerConcurrency(a, b string, timeout time.Duration) (string,error) {
    select {
        case <- ping(a):
            return a, nil
        case <- ping(b):
            return b, nil
        case <- time.After(timeout):
            return "", fmt.Errorf("time out")
    }
}

func ping(a string) chan struct{} {
    ch := make(chan struct{})

    go func() {
        http.Get(a)
        close(ch)
    }()

    return ch
}

func measureTime(a string) time.Duration {
    start := time.Now()
    http.Get(a)
    return time.Since(start)
}
