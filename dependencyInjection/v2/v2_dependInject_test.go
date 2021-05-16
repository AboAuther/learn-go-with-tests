package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "chj")
	got := buffer.String()
	want := "hello,chj"
	if got != want {
		t.Errorf("got:%q,want:%q", got, want)
	}
}
