package binarysearch

func BinarySearch(numbers []int, target int) int {
	low := 0
	high := len(numbers)

	for low < high {
		mid := low + (high-low)/2

		value := numbers[mid]
		if value == target {
			return mid
		} else if value < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return -1
}
