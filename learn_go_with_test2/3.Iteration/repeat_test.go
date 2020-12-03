package iteration

import (
    "fmt"
    "testing"
)

func TestRepeat(t *testing.T) {
    commonAssert := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got (%s) want (%s)", got, want)
        }
    }

    t.Run("test1", func(t *testing.T) {
        got := Repeat("a", 10)
        want := "aaaaaaaaaa"
        commonAssert(t, got, want)
    })
}

func BenchmarkRepeat(b *testing.B) {
    for i:=0; i<b.N; i++ {
        Repeat("a", 10)
    }
}

func ExampleRepeat() {
    r := Repeat("a", 10)
    fmt.Println(r)

    r = Repeat("a", 5)
    fmt.Println(r)

    r = Repeat("x", 5)
    fmt.Println(r)

    //Output:
    // aaaaaaaaaa
    // aaaaa
    // xxxxx
}
