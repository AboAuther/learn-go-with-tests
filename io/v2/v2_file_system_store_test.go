package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		store := FileSystemPlayerStore{database}
		got := store.GetLeague()
		want := []Player{
			{"lrg", 10},
			{"chj", 20},
		}
		assertLeague(t, got, want)
		// read again
		got = store.GetLeague()
		assertLeague(t, got, want)
	})
	t.Run("/get player score", func(t *testing.T) {
		database := strings.NewReader(`[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		store := FileSystemPlayerStore{database}
		got := store.GetPlayerScore("lrg")
		want := 10
		assertScoreEquals(t, got, want)
	})
}
func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
