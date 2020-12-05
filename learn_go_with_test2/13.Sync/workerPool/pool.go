package main

import (
    "fmt"
    "os"
    "strconv"
    "sync"
)

type Request struct {
    clientid   int
    requestNum int
}

type Response struct {
    req        Request
    answer     int
}

var (
    chanSize = 10
    reqChan = make(chan Request, chanSize) 
    resChan = make(chan Response, chanSize)
)

var wg sync.WaitGroup

func create(jobs int) {
    for i:=0; i<jobs; i++ {
        reqChan <- Request{i, i}
    }
    close(reqChan)
}

func createWorker(numWorker int) {

    for i:=0; i<numWorker; i++ {
        wg.Add(1)
        go worker(i)
    }
    wg.Wait()

    close(resChan)
}

func worker(workerid int) {
    defer wg.Done()

    for req := range reqChan {
        myResponse := Response {
            Request{ workerid, req.requestNum },
            req.requestNum * req.requestNum,
        }

        resChan <- myResponse
    }
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("jobs worker")
        os.Exit(1)
    }

    nJobs, _ := strconv.Atoi(os.Args[1])
    nWorker, _ := strconv.Atoi(os.Args[2])

    go create(nJobs)

    finished := make(chan bool)

    go func() {
        for v := range resChan {
            fmt.Printf("%v : %d\n", v.req, v.answer)
        }
        finished <- true
    }()

    createWorker(nWorker)

    <- finished
    fmt.Println("Exit...")
}
