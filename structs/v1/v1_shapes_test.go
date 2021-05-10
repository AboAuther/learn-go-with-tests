package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(10.0, 20.1)
	want := 60.2
	if got != want {
		t.Errorf("got %.2f ,wnat %.2f", got, want)
	}
}
