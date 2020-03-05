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
type Triangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}
func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
func (r Rectangle) Area() float64 {
	return (r.Height * r.Width)
}

func (C Circle) Perimeter() float64 {
	return (C.Radius) * 2.0 * math.Pi
}

func (C Circle) Area() float64 {
	return (C.Radius) * C.Radius * math.Pi
}
