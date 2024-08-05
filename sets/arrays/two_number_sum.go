package arrays

import "sort"

// Time complexity: O(n^2) | Space complexity: O(1)
func TwoNumberSumUsingBruteForce(array []int, target int) []int {
	counter := 0

	if len(array) > 1 {
		for counter <= len(array)-1 {
			for i := 0; i < len(array)-1; i++ {
				currentNumber := array[counter]

				if currentNumber != array[i] && currentNumber+array[i] == target {
					return []int{currentNumber, array[i]}
				}
			}

			counter++
		}
	}

	return []int{}
}

// Time complexity: O(n) | Space complexity: O(n)
func TwoNumberSumUsingHashing(array []int, target int) []int {
	results := make(map[int]bool)

	for _, num := range array {
		potentialMatch := target - num

		if results[num] {
			return []int{potentialMatch, num}
		}

		results[potentialMatch] = true

	}

	return []int{}
}

// Time complexity: O(n log n) | Space complexity: O(1)
func TwoNumberSumUsingSortingAndTwoPointers(array []int, target int) []int {
	sort.Ints(array)

	left := 0
	right := len(array) - 1

	if len(array) > 1 {
		for array[left]+array[right] != target {
			if array[left] != array[right] {
				if array[left]+array[right] < target {
					left++
				} else {
					right--
				}
			} else {
				return []int{}
			}
		}

		return []int{array[left], array[right]}
	}

	return []int{}
}
