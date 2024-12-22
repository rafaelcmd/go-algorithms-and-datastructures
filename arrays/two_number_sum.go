package arrays

import "sort"

// Time complexity: O(n^2) | Space complexity: O(1)
func TwoNumberSumUsingBruteForce(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		firstNum := array[i]
		for j := i + 1; j < len(array); j++ {
			secondNum := array[j]
			if firstNum+secondNum == target {
				return []int{firstNum, secondNum}
			}
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
		} else {
			results[potentialMatch] = true
		}
	}

	return []int{}
}

// Time complexity: O(n log n) | Space complexity: O(1)
func TwoNumberSumUsingSortingAndTwoPointers(array []int, target int) []int {
	sort.Ints(array)
	left := 0
	right := len(array) - 1

	for left < right {
		currentSum := array[left] + array[right]

		if currentSum == target {
			return []int{array[left], array[right]}
		} else if currentSum < target {
			left++
		} else if currentSum > target {
			right--
		}
	}

	return []int{}
}
