package bubblesort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	t.Run("sort the array with bubble sort", func(t *testing.T) {
		array := []int{1, 2, 7, 4, 3}
		got := BubbleSort(array)
		want := []int{1, 2, 3, 4, 7}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v, given %v", got, want, array)
		}
	})
}

func ExampleBubbleSort() {
	array := []int{1, 5, 8, 3, 4, 2}
	result := BubbleSort(array)
	fmt.Println(result)
	// Output: []int{1 2 3 4 5 8}
}
