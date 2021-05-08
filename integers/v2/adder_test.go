package integers

import "testing"
import "fmt"

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	want := 4
	if sum != want {
		t.Errorf("sum: %d,wnat: %d", sum, want)
	}

}
func ExampleAdd() {
	sum := Add(1, 5)

	fmt.Println(sum)
	//Output: 6

}
