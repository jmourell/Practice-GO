package structs

import "math"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}
type Circle struct {
	Radius float64
}

func (R Rectangle) Perimeter() float64 {
	return (R.Height + R.Width) * 2
}

func (R Rectangle) Area() float64 {
	return (R.Height * R.Width)
}

func (C Circle) Perimeter() float64 {
	return (C.Radius) * 2.0 * math.Pi
}

func (C Circle) Area() float64 {
	return (C.Radius) * C.Radius * math.Pi
}
