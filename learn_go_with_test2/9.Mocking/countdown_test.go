package main

import (
    "bytes"
    "reflect"
    "testing"
    "time"
)

func TestCountdown(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        buf := &bytes.Buffer{}

        Countdown(buf)

        got := buf.String()
        want := `3
2
1
Go!`

        assertStrings(t, got, want)
    })

    t.Run("default sleeper", func(t *testing.T) {
        buf := &bytes.Buffer{}
        defaultSleeper := &DefaultSleeper{}

        CountdownSleeper(buf, defaultSleeper)

        got := buf.String()
        want := `3
2
1
Go!`

        assertStrings(t, got, want)
    })

    t.Run("spy sleeper", func(t *testing.T) {
        buf := &bytes.Buffer{}
        spySleeper := &SpySleeper{}

        CountdownSleeper(buf, spySleeper)
        got := buf.String()
        want := `3
2
1
Go!`

        assertStrings(t, got, want)
        if spySleeper.Calls != 4 {
            t.Errorf("calls is not 4")
        }
    })

    t.Run("operation spy", func(t *testing.T) {
        op := &CountdownOperationSpy{}

        CountdownSleeper(op, op)

        want := []string{
            "sleep",
            "write",
            "sleep",
            "write",
            "sleep",
            "write",
            "sleep",
            "write",
        }

        if !reflect.DeepEqual(op.Calls, want) {
            t.Errorf("got (%v) want (%v)", op.Calls, want)
        }
    })

    t.Run("configurable sleeper", func(t *testing.T) {
        sleepTime := 5 * time.Second

        spyTime := &SpyTime{}
        sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
        sleeper.Sleep()

        if spyTime.durationSlept != sleepTime {
            t.Errorf("must sleep (%v) but sleep (%v)", sleepTime, spyTime.durationSlept)
        }
    })
}

func assertStrings(t *testing.T, got, want string) {
    t.Helper()

    if got != want {
        t.Errorf("got (%s) want (%s)", got, want)
    }
}
