package main

import "fmt"

func main() {
    done := make(chan struct{})
    num := 0

    go func() {
        num = 5
        done <- struct{}{}
    }()

    <- done
    <- done

    fmt.Println(num)
}
