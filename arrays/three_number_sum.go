package arrays

import "sort"

// Time: O(N²) | Space: O(N²)
func ThreeNumberSumWithSortingAndTwoPointers(array []int, target int) [][]int {
	result := [][]int{}

	sort.Ints(array)

	for current := 0; current < len(array); current++ {
		left := current + 1
		right := len(array) - 1

		for left < right {
			if array[current]+array[left]+array[right] == target {
				newArray := []int{array[current], array[left], array[right]}
				result = append(result, newArray)
				left++
				right--
			} else if array[current]+array[left]+array[right] < target {
				left++
			} else if array[current]+array[left]+array[right] > target {
				right--
			} else {
				current++
			}
		}
	}

	return result
}

// Time: O(N²) | Space: O(N²)
func ThreeNumberSumWithSortingAndTwoPointersOptimized1(array []int, target int) [][]int {
	result := [][]int{}

	if len(array) < 3 {
		return result
	}

	sort.Ints(array)

	for current := 0; current < len(array)-2; current++ {
		if array[current]+array[current+1]+array[current+2] > target {
			break
		}

		if current > 0 && array[current] == array[current-1] {
			continue
		}

		left := current + 1
		right := len(array) - 1

		for left < right {
			sum := array[current] + array[left] + array[right]

			if sum == target {
				result = append(result, []int{array[current], array[left], array[right]})
				left++
				right--

				for left < right && array[left] == array[left-1] {
					left++
				}

				for left < right && array[right] == array[right+1] {
					right--
				}
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return result
}

// Time: O(N²) | Space: O(N²)
func ThreeNumberSumWithSortingAndTwoPointersOptimized2(array []int, target int) [][]int {
	result := [][]int{}

	sort.Ints(array)

	for i := 0; i < len(array)-2; i++ {
		left := i + 1
		right := len(array) - 1

		for left < right {
			currentSum := array[i] + array[left] + array[right]

			if currentSum == target {
				result = append(result, []int{array[i], array[left], array[right]})
				left += 1
				right -= 1
			} else if currentSum < target {
				left += 1
			} else if currentSum > target {
				right -= 1
			}
		}
	}

	return result
}
