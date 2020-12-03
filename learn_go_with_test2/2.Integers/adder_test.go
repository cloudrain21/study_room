package integers

import (
    "fmt"
    "testing"
)

func TestAdder(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        got := Add(2,2)
        want := 4
        if got != want {
            t.Errorf("got %d want %d", got, want)
        }
    })
}

func ExampleAdd() {
    sum := Add(2,2)
    fmt.Println(sum)
    //Output:
    // 4
}
