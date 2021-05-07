package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("lrg", "")
		want := "hello,lrg"
		assertCorrectMessage(got, want)
	})
	t.Run("say hello world when an empty string is suplied", func(t *testing.T) {
		got := Hello("", "")
		want := "hello,world"
		assertCorrectMessage(got, want)
	})
	t.Run("say hello in Spanish", func(t *testing.T) {
		got := Hello("chj", spanish)
		want := "Hola,chj"
		assertCorrectMessage(got, want)
	})
	t.Run("saying hello in French", func(t *testing.T) {
		got := Hello("lrg", french)
		want := "Bonjour,lrg"
		assertCorrectMessage(got, want)
	})
}
