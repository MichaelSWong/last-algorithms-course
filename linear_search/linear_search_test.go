package linearsearch

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	t.Run("finds the item linearly", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		target := 2
		got := LinearSearch(numbers, target)
		want := true

		if got != want {
			t.Errorf("got %v want %v given %v", got, want, numbers)
		}
	})

	t.Run("doesn't find the item linearly", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6}
		target := 10
		got := LinearSearch(numbers, target)
		want := false

		if got != want {
			t.Errorf("got %v want %v given %v", got, want, numbers)
		}
	})
}

func ExampleLinearSearch() {
	numbers := []int{1, 5, 6, 9}
	target := 27
	result := LinearSearch(numbers, target)
	fmt.Println(result)
	// Output: false
}
