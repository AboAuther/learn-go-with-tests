package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create temp file %v", err)
	}
	_, _ = tmpfile.Write([]byte(initialData))
	removeFile := func() {
		_ = tmpfile.Close()
		_ = os.Remove(tmpfile.Name())
	}
	return tmpfile, removeFile
}
func TestFileSystemStore(t *testing.T) {
	t.Run(" league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
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
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		got := store.GetPlayerScore("lrg")
		want := 10
		assertScoreEquals(t, got, want)
	})
	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDataBase := createTempFile(t, `[
		{"Name":"lrg","Wins":10},
		{"Name":"chj","Wins":20}]`)
		defer cleanDataBase()
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
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
		store, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
		store.RecordWin("lyl")
		got := store.GetPlayerScore("lyl")
		want := 1
		assertScoreEquals(t, got, want)
	})
	t.Run("works with an empty file", func(t *testing.T) {
		database, cleanDataBase := createTempFile(t, "")
		defer cleanDataBase()
		_, err := NewFileSystemPlayerStore(database)
		assertNoError(t, err)
	})
}
func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't except an error but get one,%v", err)
	}
}
