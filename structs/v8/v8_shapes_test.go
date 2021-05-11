package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
func TestArea(t *testing.T) {
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "circle", shape: Circle{10.0}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{12.0, 6.0}, hasArea: 36.0},
	}
	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
