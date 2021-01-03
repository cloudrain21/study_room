package main

import (
    "fmt"
    "time"
)

func retryFunc(doit func(a []string) bool, args ...string) bool {

	for i := 0; i < 5; i++ {
		if doit(args) {
			return true
		}
		time.Sleep(1 * time.Second)
	}

	return false
}

func main() {
    myFunc := func(args []string) bool {
        for _, arg := range args {
            fmt.Print(arg, " ")
        }
        fmt.Println()
        return false
    }

    retryFunc(myFunc, "aa", "bb")
}
