package main

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

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Width + rec.Height)
}
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
