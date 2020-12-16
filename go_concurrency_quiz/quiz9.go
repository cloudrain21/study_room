package main

import (
    "fmt"
    "time"
)

func main() {
    ch := make(chan int)

    go func() {
        for i:=0; i<5; i++ {
            ch <- i
            time.Sleep(1 * time.Second)
        }
        close(ch)
    }()

    for {
        fmt.Println(<- ch)
    }

    fmt.Println("done")
}
