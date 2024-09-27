package main

import "testing"

func TestSum(t *testing.T) {
	nums := [5]int{4, 4, 5, 5, 1}

	got := Sum(nums)
	want := 19

	if got != want {
		t.Errorf("Got: %q, Want: %q, Given: %v", got, want, nums)
	}

}
