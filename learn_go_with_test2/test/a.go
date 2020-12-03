package main

import "fmt"

func Array() {
    a := []int{1,2,3,4,5,6,7}
    b := [...]int{1,2,3,4,5,6,7}
    c := [7]int{1,2,3,4,5,6,7}

    fmt.Println(len(a), cap(a), len(b), cap(b), len(c), cap(c))
    fmt.Printf("(%v) (%v) (%v)\n", a, b, c)

    a = append(a, 8)
    //b = append(a, 8) // error
    //c = append(a, 8) // error

    fmt.Println(len(a), cap(a), len(b), cap(b), len(c), cap(c))
    fmt.Printf("(%v) (%v) (%v)\n", a, b, c)
}

func String1() {
    s := "abcde"

    for _, c := range s {
        fmt.Println(string(c))
    }
}

func main() {
    Array()
    String1()
}
