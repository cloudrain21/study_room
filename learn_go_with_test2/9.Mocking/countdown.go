package main

import (
    "fmt"
    "io"
    "os"
    "time"
)


// use normal sleeper
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

// sleeper interface
type Sleeper interface {
    Sleep()
}


// default sleeper
type DefaultSleeper struct {}

func (d *DefaultSleeper)Sleep() {
    time.Sleep(1 * time.Second)
}

// spy sleeper
type SpySleeper struct {
    Calls int
}

func (s *SpySleeper) Sleep() {
    s.Calls++
}

// spy countdown operations
// CountdownOperationSpy implements Sleeper and io.Writer
type CountdownOperationSpy struct {
    Calls  []string
}

func (s *CountdownOperationSpy)Sleep() {
    s.Calls = append(s.Calls, "sleep")
}

func (s *CountdownOperationSpy)Write(w []byte) (int, error) {
    s.Calls = append(s.Calls, "write")
    return 0, nil
}

// configurable sleeper
type ConfigurableSleeper struct {
    duration time.Duration
    sleep func(time.Duration)
}

func (c *ConfigurableSleeper)Sleep() {
    c.sleep(c.duration)
}

type SpyTime struct {
    durationSlept time.Duration
}

func (s *SpyTime)Sleep(duration time.Duration) {
    s.durationSlept = duration
}

func CountdownSleeper(w io.Writer, s Sleeper) {
    for i:=countdownStart; i>0; i-- {
        s.Sleep()
        fmt.Fprintln(w, i)
    }
    s.Sleep()
    fmt.Fprintf(w, finalWord)
}

func main() {
    Countdown(os.Stdout)
    CountdownSleeper(os.Stdout, &DefaultSleeper{})
    CountdownSleeper(os.Stdout, &SpySleeper{})

    sleeper := &ConfigurableSleeper{1*time.Second, time.Sleep}
    CountdownSleeper(os.Stdout, sleeper)

    spyTime := &SpyTime{}
    sleeper = &ConfigurableSleeper{77*time.Second, spyTime.Sleep}
    CountdownSleeper(os.Stdout, sleeper)
    fmt.Println(spyTime.durationSlept)
}
