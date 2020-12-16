package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)
    i := 0

    go func() {
        for {
            i++; ch <- i
            time.Sleep(1 * time.Second)
        }
    }()

    for {
        fmt.Println(<- ch)
    }
}
