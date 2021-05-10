package practice

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	want := "aaaaa"
	if repeated != want {
		t.Errorf("repeated:%q,want:%q", repeated, want)
	}
}
