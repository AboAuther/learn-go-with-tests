package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(2.0, 3.0)
	want := 10.0
	if got != want {
		t.Errorf("got %.2f,want %.2f", got, want)
	}
}
func TestAreas(t *testing.T) {
	got := Areas(2.0, 3.4)
	want := 6.8
	if got != want {
		t.Errorf("got %.2f,want %.2f", got, want)
	}
}
