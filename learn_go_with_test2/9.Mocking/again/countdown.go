package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

const finalWord = "Go!"
const countdownStart = 3

func Countdown(w io.Writer) {
	for i:=countdownStart; i>0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprintf(w, finalWord)
}

func CountdownWithSleeper(w io.Writer, sleeper Sleeper) {
	for i:=3; i>0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
    sleeper.Sleep()
	fmt.Fprintf(w, finalWord)
}

type DefaultSleeper struct{}

func (d *DefaultSleeper)Sleep() {
	time.Sleep(1 * time.Second)
}

const (
    write = "write"
    sleep = "sleep"
)

type CountdownOperationSpy struct {
    Calls []string
}

func (c *CountdownOperationSpy)Sleep() {
    c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationSpy)Write(b []byte) (n int, e error) {
    c.Calls = append(c.Calls, write)
    return
}


func main() {
	CountdownWithSleeper(os.Stdout, &DefaultSleeper{})
}
