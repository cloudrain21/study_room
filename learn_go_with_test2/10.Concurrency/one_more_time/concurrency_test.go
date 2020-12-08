package concurrency

import (
    "reflect"
    "testing"
    "time"
)

func mockWebsiteChecker(url string) bool {
    if url == "want://furhrterwe.geds" {
        return false
    }
    return true
}

func slowWebsiteChecker(_ string) bool {
    time.Sleep(20 * time.Millisecond)
    return true
}

func TestCheckWebsites(t *testing.T) {
    websites := []string {
        "http://google.com",
        "http://blog.gypsydave5.com",
        "want://furhrterwe.geds",
    }

    want := map[string]bool {
        "http://google.com" : true,
        "http://blog.gypsydave5.com" : true,
        "want://furhrterwe.geds" : false,
    }

    got := CheckWebsites(mockWebsiteChecker, websites)

    if !reflect.DeepEqual(got, want) {
        t.Errorf("got %v want %v", got, want)
    }
}

func BenchmarkCheckWebsites(b *testing.B) {
    urls := make([]string, 100)
    for i:=0; i<len(urls); i++ {
        urls[i] = "a url"
    }

    for i:=0; i<b.N; i++ {
        CheckWebsites(slowWebsiteChecker, urls)
    }
}
