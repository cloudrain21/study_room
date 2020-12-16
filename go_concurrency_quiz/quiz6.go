package main

import (
    "fmt"
    "sync"
    "time"
)

var count = 0

func add(c int, wg *sync.WaitGroup) {
    time.Sleep(1 * time.Second)

    defer wg.Done()

    for i:=0; i<c; i++ {
        count++
    }
}

func main() {
    var wg sync.WaitGroup

    wg.Add(4)

    go add(10000000,&wg)
    go add(10000000,&wg)
    go add(10000000,&wg)
    go add(10000000,&wg)

    wg.Wait()

    fmt.Println(count)
}
