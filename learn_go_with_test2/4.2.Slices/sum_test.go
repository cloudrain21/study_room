package main

import (
    "fmt"
    "reflect"
    "testing"
)

func TestSum(t *testing.T) {
    commonAssert := func(t *testing.T, got, want int) {
        t.Helper()

        if got != want {
            t.Errorf("got (%d) want (%d)", got, want)
        }
    }

    t.Run("test1", func(t *testing.T) {
        n := []int{1,2,3,4,5}
        got := Sum(n)
        want := 15

        commonAssert(t, got, want)
    })

    t.Run("test2", func(t *testing.T) {
        n := []int{1,2,3}
        got := Sum(n)
        want := 6

        commonAssert(t, got, want)
    })
}

func TestSumAll(t *testing.T) {
    t.Run("test3", func(t *testing.T) {
        got := SumAll1([]int{1,2,3}, []int{1,2,3,4,5})
        want := []int{6,15}

        //if got != want {
        //    t.Errorf("got (%v) want(%v)", got, want)
        //}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got (%v) want(%v)", got, want)
        }
    })

    t.Run("be careful to use reflect.DeepEqual", func(t *testing.T) {
        got := SumAll2([]int{1,2,3}, []int{1,2,3,4,5})
        //want := "silly"  // silly way
        want := []int{6,15}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got (%v) want(%v)", got, want)
        }
    })

    t.Run("another way to sum all", func(t *testing.T) {
        got := SumAll2([]int{1,2,3}, []int{1,2,3,4,5})
        want := []int{6,15}

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got (%v) want (%v)", got, want)
        }
    })
}

func TestSumAllTails(t *testing.T) {
    commonAssert := func(t *testing.T, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got (%v) want (%v)", got, want)
        }
    }

    t.Run("sum of tails", func(t *testing.T) {
        got := SumAllTails([]int{0,2,3}, []int{1,2,9})
        want := []int{5,11}

        commonAssert(t, got, want)
    })

    t.Run("empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{1,2,9})
        want := []int{0,11}

        commonAssert(t, got, want)
    })
}

func ExampleSum() {
    n := []int{1,2,3,4,5}
    fmt.Println(Sum(n))

    // Output:
    // 15
}
