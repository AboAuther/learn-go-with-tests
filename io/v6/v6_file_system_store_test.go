package main

import (
	"io"
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(t *testing.T, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}
func TestFileSystemStore(t *testing.T) {
	t.Run(" league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		defer cleanDatabase()
		store := NewFileSystemPlayerStore(database)
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
	t.Run(" get player score", func(t *testing.T) {
		database, cleanDataBase := createTempFile(t, `[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		defer cleanDataBase()
		store := NewFileSystemPlayerStore(database)
		got := store.GetPlayerScore("lrg")
		want := 10
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDataBase := createTempFile(t, `[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		defer cleanDataBase()
		store := NewFileSystemPlayerStore(database)
		store.RecordWin("chj")
		got := store.GetPlayerScore("chj")
		want := 21
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDataBase := createTempFile(t, `[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		defer cleanDataBase()
		store := NewFileSystemPlayerStore(database)
		store.RecordWin("lyl")
		got := store.GetPlayerScore("lyl")
		want := 1
		assertScoreEquals(t, got, want)
	})
}
func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
