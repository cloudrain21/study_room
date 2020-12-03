package main

import "testing"

func TestHello(t *testing.T) {
    assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run( "test1", func(t *testing.T) {
        got := Hello("Chris", "English")
        want := "Hello, Chris"
        assertCorrectMessage(t, got, want)
    })

    t.Run( "test2", func(t *testing.T) {
        got := Hello("", "English")
        want := "Hello, World"
        assertCorrectMessage(t, got, want)
    })

    t.Run( "test3", func(t *testing.T) {
        got := Hello("Elodie", "Spanish")
        want := "Hola, Elodie"
        assertCorrectMessage(t, got, want)
    })
}
