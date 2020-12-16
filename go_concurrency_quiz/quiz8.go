package main

import "fmt"

func main() {
    done := make(chan struct{})

    go func() {
        done <- struct{}{}
        close(done)
    }()

    <- done

    fmt.Println("done")
}
