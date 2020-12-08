package clockface

import (
    "math"
    "time"
    "testing"
)

func TestSecondHandAtMidnight(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

        want := Point{X: 150, Y:150 - 90}
        got := SecondHand(tm)

        if got != want {
            t.Errorf("got %v want %v", got, want)
        }
    })

    //t.Run("test2", func(t *testing.T) {
    //    tm := time.Date(1337, time.January, 1, 0, 30, 0, 0, time.UTC)

    //    want := Point{X: 150, Y:150 + 90}
    //    got := SecondHand(tm)

    //    if got != want {
    //        t.Errorf("got %v want %v", got, want)
    //    }
    //})
}

func TestSecondsInRadians(t *testing.T) {
    t.Run("test thirty second", func(t *testing.T) {
        thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30,0, time.UTC)

        want := math.Pi
        got := secondsInRadians(thirtySeconds)

        if want != got {
            t.Errorf("got %v want %v", got, want)
        }
    })

    cases := []struct{
        time time.Time
        angle float64
    } {
        {simpleTime(0, 0, 30), math.Pi},
        {simpleTime(0, 0, 0),  0},
        {simpleTime(0, 0, 45), (math.Pi /2) * 3},
        {simpleTime(0, 0, 7), (math.Pi /30) * 7},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            want := c.angle
            got := secondsInRadians(c.time)

            if want != got {
                t.Errorf("got %v want %v", got, want)
            }
        })
    }
}

func TestSecondHandVector(t *testing.T) {
    cases := []struct {
        time time.Time
        point Point
    } {
        {simpleTime(0,0,30), Point{0,-1}},
        {simpleTime(0,0,45), Point{-1,0}},
    }

    for _, c := range cases {
        t.Run(testName(c.time), func(t *testing.T) {
            want := c.point
            got := secondHandPoint(c.time)

            if roughlyEqual(got, want) {
                t.Errorf("got %v want %v", got, want)
            }
        })
    }
}
