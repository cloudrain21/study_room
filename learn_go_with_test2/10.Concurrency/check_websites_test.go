package concurrency

import (
    "reflect"
    "testing"
    "time"
)

func mockWebsiteChecker(url string) bool {
    if url == "waat://furhurterwe.geds" {
        return false
    }
    return true
}

func slowStubWebsiteChecker(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

func TestCheckWebsites(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        websites := []string {
            "http://cloudrain21.com",
            "https://abc.com",
            "waat://furhurterwe.geds",
        }

        want := map[string]bool {
            "http://cloudrain21.com": true,
            "https://abc.com": true,
            "waat://furhurterwe.geds": false,
        }

        got := CheckWebsites(mockWebsiteChecker, websites)

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got (%v) want (%v)", got, want)
        }
    })
}

func BenchmarkCheckWebsites(b *testing.B) {
    urls := make([]string, 100)
    for i:=0; i<100; i++ {
        urls[i] = "a url"
    }

    for i:=0; i<b.N; i++ {
        CheckWebsites(slowStubWebsiteChecker, urls)
    }
}
