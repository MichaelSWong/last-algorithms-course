package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	t.Run("finds the index of the target binaryly", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target := 2
		got := BinarySearch(numbers, target)
		want := 1

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
	t.Run("finds the index of the target binaryly larger", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target := 8
		got := BinarySearch(numbers, target)
		want := 7

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("doesn't find the index of the target binaryly", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		target := 15
		got := BinarySearch(numbers, target)
		want := -1

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

func ExampleBinarySearch() {
	numbers := []int{1, 2, 3}
	target := 2
	result := BinarySearch(numbers, target)
	fmt.Println(result)
	// Output: 1
}
