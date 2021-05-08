package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("to a person", func(t *testing.T) {
		got := Hello("lrg", "")
		want := "hello,lrg"
		assertCorrectMessage(t, got, want)
	})
	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "hello,world"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("chj", spanish)
		want := "Hola,chj"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("lrg", french)
		want := "Bonjour,lrg"
		assertCorrectMessage(t, got, want)
	})
}
