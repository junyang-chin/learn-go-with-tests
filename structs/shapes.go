package structs

import "math"

func Perimeter(shape Rectangle) float64 {
	return 2 * (shape.Width + shape.Height)
}

func Area(shape Rectangle) float64 {
	return shape.Width * shape.Height
}

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

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
