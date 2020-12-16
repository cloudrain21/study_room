package main

import (
    . "fmt"
    "time"
)

func main() {
    ch := make(chan int)

    go func() {
        Println("before")
        ch <- 1
        Println("after")
    }()

    for {
        time.Sleep(1 * time.Second)
    }
}
