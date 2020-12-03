package main

import (
    "testing"
)

func TestSearch(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        dic := map[string]string{"test":"this is test value"}

        got := Search(dic, "test")
        want := "this is test value"

        assertStrings(t, got, want)
    })

    t.Run("mysearch test1", func(t *testing.T) {
        d := Dictionary{"test":"this is test value"}

        got,_ := d.Search("test")
        want := "this is test value"

        assertStrings(t, got, want)
    })

    t.Run("known key", func(t *testing.T) {
        dic := Dictionary{"known":"this is test value"}

        got, _ := dic.Search("known")
        want := "this is test value"

        assertStrings(t, got, want)
    })

    t.Run("unknown key", func(t *testing.T) {
        dic := Dictionary{"known":"this is test value"}

        _, err := dic.Search("unknown")
        want := "this is unknown key"

        if err == nil {
            t.Fatal("error must not be nil")
        }

        assertStrings(t, err.Error(), want)
    })

    t.Run("unknown key 2", func(t *testing.T) {
        dic := Dictionary{"test":"this is test value"}
        _, err := dic.Search("unknown")

        assertError(t, err, ErrNotFound)
    })

    t.Run("add new word", func(t *testing.T) {
        dic := Dictionary{"key1":"value1"}

        var err error
        err = dic.Add("key2","value2")
        if err != nil {
            t.Fatal("aaaaaaa", err)
        }

        got, err := dic.Search("key2")
        if err != nil {
            t.Fatal("key2 fatal : ", err)
        }

        want := "value2"

        assertStrings(t, got, want)
    })

    t.Run("already exists", func(t *testing.T) {
        dic := Dictionary{"key1":"value1"}

        err := dic.Add("key1", "value1")

        assertError(t, err, ErrAlreadyExist)
    })

    t.Run("update", func(t *testing.T) {
        dic := Dictionary{"key1":"value1"}

        dic.Update("key1", "value111")

        got, err := dic.Search("key1")
        if err != nil {
            t.Fatal("error must be nil")
        }
        want := "value111"

        assertStrings(t, got, want)
    })

    t.Run("update not found", func(t *testing.T) {
        dic := Dictionary{"key1":"value1"}

        err := dic.Update("key_not_found", "value111")

        assertError(t, err, ErrNotFound)
    })

    t.Run("update okay", func(t *testing.T) {
        dic := Dictionary{"key1":"value1"}

        err := dic.Update("key1", "value111")

        assertError(t, err, nil)

        got, err := dic.Search("key1")
        want := "value111"

        assertError(t, err, nil)
        assertStrings(t, got, want)
    })

    t.Run("delete", func(t *testing.T) {
        dic := Dictionary{"key1":"value1"}

        dic.Delete("key1")

        got, err := dic.Search("key1")

        assertError(t, err, ErrNotFound)
        assertStrings(t, got, "")
    })
}

func assertStrings(t *testing.T, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got (%s) want (%s)", got, want)
    }
}

func assertError(t *testing.T, got, want error) {
    t.Helper()

    if got != want {
        t.Errorf("got (%s) want (%s)", got, want)
    }
}
