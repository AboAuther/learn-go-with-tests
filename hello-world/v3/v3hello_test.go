package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("lrg")
	want := "hello,lrg"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
