package main

import (
    "fmt"
    "time"
)

func pingpong(name string, ch chan int) {
    for {
        i := <- ch
        fmt.Println(name, i)

        ch <- i
        time.Sleep(1 * time.Second)
    }
}

func main() {
    ch := make(chan int)

    go pingpong("player1", ch)
    go pingpong("player2", ch)
}
