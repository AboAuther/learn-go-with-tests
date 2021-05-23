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
}
