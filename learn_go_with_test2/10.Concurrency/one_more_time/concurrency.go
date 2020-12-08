package concurrency

type WebsiteChecker func(string) bool

type Result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    result := make(map[string]bool)
    rChan := make(chan Result)

    for _, url := range urls {
        go func(u string) {
            rChan <- Result{u, wc(u)}
        }(url)
    }

    for i:=0; i<len(urls); i++ {
        r := <- rChan
        result[r.string] = r.bool
    }

    return result
}
