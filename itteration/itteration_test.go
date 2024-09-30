package itteration

import (
	"fmt"
	"testing"
)

func TestItteration(t *testing.T) {

	got := RepeatNTimes("a", 5)
	want := "aaaaa"

	if got != want {
		t.Errorf("Got: %q, Want: %q", got, want)
	}
}

func ExampleRepeat() {
	got := RepeatNTimes("R", 4)
	fmt.Println(got)
	// Output: RRRR
}

func BenchmarkRepeatNTimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatNTimes("a", 10)
	}
}
