package main

import (
    "fmt"
    "time"
)

func pingpong(name string, ch chan int) {
    for {
        i := <- ch
        fmt.Println(name, ":", i)

        i++
        ch <- i
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch := make(chan int)

    go pingpong("player1", ch)
    go pingpong("player2", ch)

    ch <- 1

    for {
        time.Sleep(1 * time.Second)
    }
}
