package mystrings

import (
    "fmt"
    "testing"
)

func TestStrings(t *testing.T) {
    commonAssertInterface := func(t *testing.T, got, want interface{}) {
        t.Helper()

        switch got.(type) {
            case int:
                if got != want {
                    t.Errorf("got (%d) want(%d)", got, want)
                }
            case string:
                if got != want {
                    t.Errorf("got (%s) want(%s)", got, want)
                }
            case bool:
                if got != want {
                    t.Errorf("got (%v) want(%v)", got, want)
                }
        }

    }

    t.Run("join", func(t *testing.T) {
        s1 := "abc"
        s2 := "def"

        got := MyJoin(s1, s2)
        want := "abcdef"

        commonAssertInterface(t, got, want)
    })

    t.Run("contains", func(t *testing.T) {
        s := "abcdefg"
        f := "cde"

        got := MyContains(s, f)
        want := true

        commonAssertInterface(t, got, want)
    })

    t.Run("index", func(t *testing.T) {
        s := "abcdefg"
        f := "cde"

        got := MyIndex(s, f)
        want := 2

        commonAssertInterface(t, got, want)
    })
}

func ExampleStrings() {
    fmt.Println(MyJoin("abc", "def"))

    // Output:
    // abcdef
}

func BenchmarkStrings(b *testing.B) {
    for i:=0; i<b.N; i++ {
        MyJoin("abc", "def")
    }
}
