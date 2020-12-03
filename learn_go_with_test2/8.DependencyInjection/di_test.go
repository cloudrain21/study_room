package main

import (
    "bytes"
    "testing"
)

func TestGreet(t *testing.T) {
    t.Run("greet", func(t *testing.T) {
        buffer := bytes.Buffer{}
        Greet(&buffer, "chris")

        got := buffer.String()
        want := "Hello, chris"

        assertStrings(t, got, want)
    })
}

func assertStrings(t *testing.T, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got (%s) want (%s)", got, want)
    }
}
