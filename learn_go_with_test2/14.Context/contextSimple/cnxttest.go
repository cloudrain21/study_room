package main

import (
    "fmt"
    "context"
    "os"
    "strconv"
    "time"
)

func f1(delay int) {
    cnxt := context.Background()

    cnxt, cancel := context.WithCancel(cnxt)
    defer cancel()

    go func() {
        time.Sleep(4 * time.Second)
        cancel()
    }()

    select {
    case <- cnxt.Done():
        fmt.Println("cancel done...")
    case r := <- time.After(time.Duration(delay) * time.Second):
        fmt.Println("f1() : ", r)
    }
}

func f2(delay int) {
    cnxt := context.Background()
    cnxt, cancel := context.WithTimeout(cnxt, time.Duration(delay) * time.Second)
    defer cancel()

    go func() {
        time.Sleep(4 * time.Second)
        cancel()
    }()

    select {
    case <- cnxt.Done():
        fmt.Println("cancel done...")
    case r := <- time.After(time.Duration(delay) * time.Second):
        fmt.Println("f2() : ", r)
    }
}

func f3(delay int) {
    cnxt := context.Background()
    deadline := time.Now().Add(time.Duration(delay) * time.Second)
    cnxt, cancel := context.WithDeadline(cnxt, deadline)
    defer cancel()

    go func() {
        time.Sleep(4 * time.Second)
        cancel()
    }()

    select {
    case <- cnxt.Done():
        fmt.Println("cancel done...")
    case r := <- time.After(time.Duration(delay) * time.Second):
        fmt.Println("f3() : ", r)
    }
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("duration must be given")
        os.Exit(1)
    }

    delay,err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("Delay : %d\n", delay)

    f1(delay)
    f2(delay)
    f3(delay)
}
