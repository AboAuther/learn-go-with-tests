package main

import "math"

// Rectangle 矩形结构
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 返回矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 返回矩形周长
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle 圆
type Circle struct {
	Radius float64
}

// Area 返回圆面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
