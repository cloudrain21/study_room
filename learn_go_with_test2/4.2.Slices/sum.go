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

func SumAll1(a ...[]int) ([]int) {
    r := make([]int, len(a))

    for i, v := range a {
        r[i] = Sum(v)
    }

    return r
}

func SumAll2(a ...[]int) []int {
    var r []int

    for _, v := range a {
        r = append(r, Sum(v))
    }

    return r
}

func SumAllTails(a ...[]int) []int {
    r := []int{}

    for _, v := range a {
        if len(v) == 0 {
            r = append(r, 0)
        } else {
            r = append(r, Sum(v[1:]))
        }
    }

    return r
}

func main() {
    n := []int{1,2,3,4,5}

    fmt.Println(Sum(n))
    fmt.Println(SumAll1([]int{1,2,3}, []int{1,2,3,4,5}))
    fmt.Println(SumAll2([]int{1,2,3}, []int{1,2,3,4,5}))
}
