package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Width float64
    Height float64
}

type Circle struct {
    Radius float64
}

type Square struct {
    Width float64
}

type Triangle struct {
    Bottom float64
    Height float64
}

func Perimeter(r Rectangle) float64 {
    return 2*(r.Width + r.Height)
}

func (r Rectangle)Area() float64 {
    return r.Width * r.Height
}

func (c Circle)Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (s Square)Area() float64 {
    return s.Width * s.Width
}

func (t Triangle)Area() float64 {
    return t.Bottom * t.Height * 0.5
}

func main() {
    fmt.Printf( "%.2f\n", Perimeter(Rectangle{10.0,10.0}) )
}
