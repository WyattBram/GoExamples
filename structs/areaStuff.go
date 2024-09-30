package structs

import "math"

type Shape interface {
	area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Width  float64
	Height float64
}

func (t Triangle) area() float64 {
	return .5 * (t.Width * t.Height)
}

func (r Rectangle) area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) area() float64 {
	return math.Pi * (c.Radius * c.Radius)
}

func Perimiter(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}
