package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "os"
    "time"
)

func random(min, max int) int {
    return rand.Intn(max - min) + min
}

func myHandler(w http.ResponseWriter, r *http.Request) {
    rand.Seed(time.Now().Unix())

    delay := random(1, 100)

    time.Sleep(time.Duration(delay) * time.Millisecond)

    fmt.Fprintf(w, "Serving : %s\n", r.URL.Path)
    fmt.Fprintf(w, "delay : %d", delay)

    fmt.Printf("serving... delay(%d)\n", delay)
}

func main() {
    var port = ":8001"

    if len(os.Args) > 1 {
        portNum := os.Args[1]
        port = ":" + portNum
    }

    fmt.Printf("port (%s)\n", port)

    http.HandleFunc("/", myHandler)

    err := http.ListenAndServe(port, nil)
    if err != nil {
        fmt.Println("listen err : ", err)
        os.Exit(10)
    }
}
