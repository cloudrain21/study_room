package main

import "fmt"

type MyInt int
type MyInt2 = int

func main() {
    var a MyInt
    var b MyInt2

    check(a)
    check(b)
}

func check(a interface{}) {
    switch a.(type) {
        case MyInt:
            fmt.Println("a is MyInt")
        //case MyInt2:
        //    fmt.Println("a is MyInt2")
        case int:
            fmt.Println("a is int")
    }
}
