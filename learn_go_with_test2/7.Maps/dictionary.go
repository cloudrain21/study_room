package main

import (
    "fmt"
)

const (
    ErrNotFound     = DictionaryErr("this is unknown key")
    ErrAlreadyExist = DictionaryErr("this key already exists")
)

// implements error interface
type DictionaryErr string

func (e DictionaryErr)Error() string {
    return string(e)
}

// simple search function
func Search(m map[string]string, key string) string {
    return m[key]
}

// user dictionary type
type Dictionary map[string]string

func (d Dictionary)Search(key string) (string,error) {
    val, ok := d[key]
    if !ok {
        return "", ErrNotFound
    }
    return val, nil
}

func (d Dictionary)Add(key, value string) error {
    _, err := d.Search(key)

    switch err {
        case ErrNotFound:
            d[key] = value
        case nil:
            return ErrAlreadyExist
        default:
            return err
    }

    return nil
}

func (d Dictionary)Update(key, value string) error {
    _, err := d.Search(key)
    if err != nil {
        return err
    }

    d[key] = value
    return nil
}

func (d Dictionary)Delete(key string) {
    delete(d,key)
}

func main() {
    dic := map[string]string{"test":"this is test value"}
    fmt.Println(Search(dic, "test"))

    d := Dictionary{"test":"this is test"}
    fmt.Println(d.Search("test"))
}
