package arrays

import (
	"slices"
	"testing"
)

func TestArraySum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := ArraySum(numbers)
		expected := 15

		if got != expected {
			t.Errorf("expected %d but got %d when given %v", expected, got, numbers)
		}
	})

	// t.Run("Collection of varying size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}

	// 	got := ArraySum(numbers)
	// 	expected := 6

	// 	if got != expected {
	// 		t.Errorf("expected %d but got %d when given %v", expected, got, numbers)
	// 	}
	// })
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	expected := []int{6, 15}

	if !slices.Equal(got, expected) {
		t.Errorf("expected %v but got %v", expected, got)
	}
}
