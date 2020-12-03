package main

import (
    "fmt"
)

func Sum(n []int) int {
    sum := 0
    for _, v := range n {
        sum += v
    }

    return sum
}

func main() {
    fmt.Println(Sum([]int{1,2,3,4,5}))
}
