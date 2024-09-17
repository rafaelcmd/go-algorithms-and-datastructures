package arrays

import "sort"

// Time: O(n * log(n)) | Space: O(n)
func NonConstructibleChange(coins []int) int {
	sort.Ints(coins)

	currentSum := 0

	for _, coin := range coins {
		if coin > currentSum+1 {
			return currentSum + 1
		}
		currentSum += coin
	}

	return currentSum + 1
}
