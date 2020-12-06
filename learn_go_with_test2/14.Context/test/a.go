package main

import "fmt"

func main() {
    s := "abcde"

    for _, v := range s {
        fmt.Printf("%c\n", v)
    }
}
