package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSumAllTails(t *testing.T) {

	t.Run("Test with valid length slices", func(t *testing.T) {
		nums1 := []int{5, 5, 1}
		nums2 := []int{5, 9, 1}
		nums3 := []int{5, 1, 1}

		got := SumAllTails(nums1, nums2, nums3)
		want := []int{6, 10, 2}

		if !slices.Equal(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

	t.Run("Test with slice len 0", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{5, 9, 1}
		nums3 := []int{5, 1, 1}

		got := SumAllTails(nums1, nums2, nums3)
		want := []int{0, 10, 2}

		if !slices.Equal(got, want) {
			t.Errorf("Got: %v, Want: %v", got, want)
		}
	})

}

func TestSum(t *testing.T) {

	t.Run("Testing with variable length", func(t *testing.T) {

		nums := []int{5, 5, 1}

		got := Sum(nums)
		want := 11

		if got != want {
			t.Errorf("Got: %q, Want: %q, Given: %v", got, want, nums)
		}
	})

	t.Run("Testing adding slices together", func(t *testing.T) {

		nums1 := []int{5, 5, 1}
		nums2 := []int{5, 9, 1}
		nums3 := []int{5, 1, 1}

		got := SumAll(nums1, nums2, nums3)
		want := []int{11, 15, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got: %q, Want: %q, Given: \n%v\n%v\n%v", got, want, nums1, nums2, nums3)
		}
	})

}
