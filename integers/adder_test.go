package integers

import "testing"
import "fmt"

func TestAdder(t *testing.T) {
	got := Add(2, 4)
	want := 6

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
