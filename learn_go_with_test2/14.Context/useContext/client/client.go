package main

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
    "sync"
    "time"
)

var (
    myUrl string
    delay int = 5
    w sync.WaitGroup
)

type myData struct {
    r *http.Response
    err error
}

func main() {
    if len(os.Args) < 3 {
        fmt.Println("need args (server url, timeout)")
        os.Exit(1)
    }

    myUrl = os.Args[1]
    delay, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Println("atoi err : ", err)
        os.Exit(2)
    }

    cnxt := context.Background()
    cnxt,cancel := context.WithTimeout(cnxt, time.Duration(delay) * time.Millisecond)
    defer cancel()

    w.Add(1)

    go connect(cnxt)

    w.Wait()
}

func connect(cnxt context.Context) {
    defer w.Done()

    dataChan := make(chan myData, 1)

    tr := &http.Transport{}
    client := &http.Client{Transport:tr}

    req, _ := http.NewRequest("GET", myUrl, nil)

    go func() {
        response, err := client.Do(req)
        if err != nil {
            fmt.Println(err)
            dataChan <- myData{nil, err}
        } else {
            fmt.Println()
            dataChan <- myData{response,nil}
        }
    }()

    select {
    case <- cnxt.Done():
        //tr.CancelRequest(req)
        <- dataChan
        fmt.Println("request was cancelled")
    case r := <- dataChan:
        err := r.err
        res := r.r
        if err != nil {
            fmt.Println("err response : ", err)
            return
        }
        defer res.Body.Close()

        httpData, err := ioutil.ReadAll(res.Body)
        if err != nil {
            fmt.Println("read body err : ", err)
            return
        }

        fmt.Printf("read data : %s\n", httpData)
    }
}
