package main

import (
    "fmt"
    "testing"
)

func TestSum(t *testing.T) {

    commonAssert := func(t *testing.T, got, want interface{}) {
        t.Helper()
        if got != want {
            switch got.(type) {
                case int:
                    fmt.Errorf("got (%d) want (%d)", got, want)
                case string:
                    fmt.Errorf("got (%q) want (%q)", got, want)
                case bool:
                    fmt.Errorf("got (%v) want (%v)", got, want)
            }
        }
    }

    t.Run("test1", func(t *testing.T) {
        n := []int{1,2,3,4,5}

        got := Sum(n)
        want := 15

        commonAssert(t, got, want)
    })
}

func ExampleSum() {
    n := []int{1,2,3,4,5}
    fmt.Println(Sum(n))

    // Output:
    // 15
}

func BenchmarkSum(b *testing.B) {
    for i:=0; i<b.N; i++ {
        Sum([]int{1,2,3,4,5})
    }
}
