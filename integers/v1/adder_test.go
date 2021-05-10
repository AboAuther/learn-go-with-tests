package integers

import "testing"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	want := 4
	if sum != want {
		t.Errorf("sum: %d,want: %d", sum, want)
	}
}
