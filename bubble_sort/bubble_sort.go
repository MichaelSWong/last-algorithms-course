package bubblesort

func BubbleSort(arr []int) []int {
	n := len(arr)
	for i := range n {
		for j := range (n) - i - 1 {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	// for i := 0; i < n-1; i++ {
	// 	for j := 0; j < n-i-1; j++ {
	// 		if arr[j] > arr[j+1] {
	// 			arr[j], arr[j+1] = arr[j+1], arr[j]
	// 		}
	// 	}
	// }
	return arr
}
