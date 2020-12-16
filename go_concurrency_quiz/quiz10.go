package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    ch := make(chan int)
    var wg sync.WaitGroup

    wg.Add(1)
    go func() {
        defer wg.Done()

        for i:=0; ;i++ {
            ch <- i
            time.Sleep(1 * time.Second)
        }
    }()

    for i:=0; i<5; i++ {
        fmt.Println(<- ch)
    }
    close(ch)

    wg.Wait()

    fmt.Println("done")
}
