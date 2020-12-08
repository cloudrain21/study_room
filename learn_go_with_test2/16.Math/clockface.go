package clockface

import (
    "math"
    "time"
)

type Point struct {
    X float64
    Y float64
}

func SecondHand(t time.Time) Point {
    return Point{150,60}
}

func secondsInRadians(t time.Time) float64 {
    return (math.Pi / (30 / float64(t.Second())))
}

func simpleTime(h, m, s int) time.Time {
    return time.Date(312, time.October, 28, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
    return t.Format("15:04:05")
}

func secondHandPoint(t time.Time) Point {
    angle := secondsInRadians(t)

    x := math.Sin(angle)
    y := math.Cos(angle)

    return Point{x,y}
}

func roughlyEqualValue(a, b float64) bool {
    return math.Abs(a-b) < 1e-7
}

func roughlyEqual(p1, p2 Point) bool {
    return roughlyEqualValue(p1.X, p2.X) &&
           roughlyEqualValue(p1.Y, p2.Y)
}
