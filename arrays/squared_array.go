package arrays

import "math"

// Time: O(n^2) | Space: O(N)
func SortedSquaredArray(array []int) []int {
	result := []int{}

	for _, v := range array {
		result = append(result, v*v)
	}

	insertionSort(result)

	return result
}

func insertionSort(array []int) {
	for i := 1; i < len(array); i++ {
		current := array[i]
		ordered := i - 1

		for ordered >= 0 && array[ordered] > current {
			array[ordered+1] = array[ordered]
			ordered--
		}

		array[ordered+1] = current
	}
}

// Time: O(n) | Space: O(N)
func SortedSquaredArrayTwoPointers(array []int) []int {
	result := make([]int, len(array))
	first := 0
	last := len(array) - 1
	counter := len(array) - 1

	for first <= last {
		if abs(array[first]) > abs(array[last]) {
			result[counter] = array[first] * array[first]
			first++
			counter--
		} else {
			result[counter] = array[last] * array[last]
			last--
			counter--
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// Time: O(n) | Space: O(N)
func SortedSquaredArrayTwoPointersOptimized(array []int) []int {
	n := len(array)
	result := make([]int, n)
	left := 0
	right := n - 1
	position := n - 1

	for left <= right {
		leftSquare := math.Abs(float64(array[left]))
		rightSquare := math.Abs(float64(array[right]))

		if leftSquare > rightSquare {
			result[position] = int(leftSquare * leftSquare)
			left++
		} else {
			result[position] = int(rightSquare * rightSquare)
			right--
		}
		position--
	}

	return result
}
