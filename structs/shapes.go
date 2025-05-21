package structs

import "math"

type Shape interface {
	Area() float64
}

// -------------- Circle
type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * (circle.Radius * circle.Radius)
}

// -------------- Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (rectangle Rectangle) Area() float64 {
	return rectangle.Width * rectangle.Height
}

// -------------- Triangle
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

// Perimeter calculates perimeter of rectangular shape
func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Width + rectangle.Height) * 2
}
