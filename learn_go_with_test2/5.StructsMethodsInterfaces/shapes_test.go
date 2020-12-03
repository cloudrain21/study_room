package main

import (
    "math"
    "testing"
)

func TestPerimeter(t *testing.T) {
    r := Rectangle{10.0, 10.0}
    got := Perimeter(r)
    want := 40.0

    if got != want {
        t.Errorf("got (%.2f) want(%.2f)", got, want)
    }
}

func TestArea(t *testing.T) {
    t.Run("area of rectangle", func(t *testing.T) {
        r := Rectangle{12.0, 6.0}
        got := r.Area()
        want := 72.0

        if got != want {
            t.Errorf("got (%.2f) want(%.2f)", got, want)
        }
    })

    t.Run("area of rectangle", func(t *testing.T) {
        r := Rectangle{12.0, 6.0}
        got := r.Area()
        want := 72.0

        if got != want {
            t.Errorf("got (%g) want(%g)", got, want)
        }
    })

    t.Run("area of circle", func(t *testing.T) {
        c := Circle{10}
        got := c.Area()
        want := 100 * math.Pi

        if got != want {
            t.Errorf("got (%g) want (%g)", got, want)
        }
    })
}

func TestAreaInterface(t *testing.T) {

    checkArea := func(t *testing.T, s Shape, want float64) {
        t.Helper()

        got := s.Area()

        if got != want {
            t.Errorf("got (%g) want (%g)", got, want)
        }
    }

    t.Run("area of rectangle", func(t *testing.T) {
        r := Rectangle{12.0, 6.0}
        want := 72.0

        checkArea(t, r, want)
    })

    t.Run("area of circle", func(t *testing.T) {
        c := Circle{10.0}
        want := 100 * math.Pi

        checkArea(t, c, want)
    })

    t.Run("area of square", func(t *testing.T) {
        s := Square{10.0}
        want := 100.0

        checkArea(t, s, want)
    })
}

func TestTableDriven(t *testing.T) {
    checkArea := func(t *testing.T, shape Shape, want float64) {
        t.Helper()

        got := shape.Area()
        if got != want {
            t.Errorf("%#v got (%g) want (%g)", shape, got, want)
        }
    }

    // table : anonymous struct
    table := []struct {
        shape Shape
        want  float64
    } {
        { shape:Rectangle{Width:12.0,   Height:6.0}, want:72.0 },
        { shape:Circle   {Radius:10.0},              want:100 * math.Pi },
        { shape:Square   {Width:10.0},               want:100.0 },
        { shape:Triangle {Bottom:10.0,  Height:6.0}, want:30.0 },
    }

    for _, v := range table {
        checkArea(t, v.shape, v.want)
    }
}
