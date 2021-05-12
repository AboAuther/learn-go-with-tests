package main

import "math"

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.width + rec.height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Height * t.Base)
}
