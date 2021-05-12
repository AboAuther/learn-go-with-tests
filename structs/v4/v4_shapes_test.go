package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{1.0, 1.0}
	got := Perimeter(rectangle)
	want := 4.0
	if got != want {
		t.Errorf("got:%.2f,want:%.2f", got, want)
	}
}
func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{2.0, 4.0}
		got := rectangle.Area()
		want := 8.0
		if got != want {
			t.Errorf("got:%.2f,want:%.2f", got, want)
		}
	})
	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		if got != want {
			t.Errorf("got:%g,want:%g", got, want)
		}
	})
}
