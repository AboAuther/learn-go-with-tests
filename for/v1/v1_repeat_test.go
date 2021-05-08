package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	want := "aaaaa"
	if repeated != want {
		t.Errorf("repeated: %s,want: %s", repeated, want)
	}
}
