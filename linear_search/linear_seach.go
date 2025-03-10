package linearsearch

import "slices"

func LinearSearch(numbers []int, target int) bool {
	return slices.Contains(numbers, target)
	// for _, number := range numbers {
	// 	if number == target {
	// 		return true
	// 	}
	// }
	//
	// return false
}
