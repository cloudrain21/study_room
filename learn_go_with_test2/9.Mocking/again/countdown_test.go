package main

import (
    "bytes"
    "reflect"
    "testing"
    "time"
)

type SpySleeper struct {
    Calls int
}

func (s *SpySleeper)Sleep() {
    s.Calls++
}

func TestCountdown(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        buf := &bytes.Buffer{}

        Countdown(buf)

        got := buf.String()
        want := `3
2
1
Go!`

        if got != want {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("spy sleeper", func(t *testing.T) {
        buf := &bytes.Buffer{}
        spySleeper := &SpySleeper{0}

        CountdownWithSleeper(buf, spySleeper)

        got := buf.String()
        want := `3
2
1
Go!`

        if got != want {
            t.Errorf("got %v want %v", got, want)
        }

        if spySleeper.Calls != 4 {
            t.Errorf("must b 4, but %d", spySleeper.Calls)
        }
    })

    t.Run("CounterdownOperationSpy", func(t *testing.T) {
        want := []string {
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
            sleep,
            write,
        }

        spy := &CountdownOperationSpy{}

        CountdownWithSleeper(spy, spy)

        if !reflect.DeepEqual(want, spy.Calls) {
            t.Errorf("got %v want %v", spy.Calls, want)
        }
    })
}

type SpyTime struct {
    durationSlept time.Duration
}

func (s *SpyTime)Sleep(duration time.Duration) {
    s.durationSlept = duration
}

type ConfigurableSleeper struct {
    duration time.Duration
    sleep func(time.Duration)
}

func (c *ConfigurableSleeper)Sleep() {
    c.sleep(c.duration)
}

func TestConfigurableSleeper(t *testing.T) {
    t.Run("configurable sleeper", func(t *testing.T) {
        sleepTime := 5 * time.Second

        spyTime := &SpyTime{}
        sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
        sleeper.Sleep()

        if spyTime.durationSlept != sleepTime {
            t.Errorf("sleepTime : %v sleptTime : %v", sleepTime, spyTime.durationSlept)
        }
    })
}