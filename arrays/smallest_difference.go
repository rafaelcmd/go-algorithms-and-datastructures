package arrays

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	smallestDifference := math.MaxInt64
	var result []int

	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			currentDifference := int(math.Abs(float64(array1[i]) - float64(array2[j])))

			if currentDifference < smallestDifference {
				smallestDifference = currentDifference
				result = []int{array1[i], array2[j]}
			}
		}
	}

	return result
}

func SmallestDifferenceWithSortingAndTwoPointers(array1, array2 []int) []int {
	smallestDifference := math.MaxInt64
	result := make([]int, 2)
	sort.Ints(array1)
	sort.Ints(array2)

	pointer1 := 0
	pointer2 := 0

	for pointer1 < len(array1) && pointer2 < len(array2) {
		if array1[pointer1] < array2[pointer2] {
			currentDifference := int(math.Abs(float64(array1[pointer1]) - float64(array2[pointer2])))
			if currentDifference < smallestDifference {
				smallestDifference = currentDifference
				result[0] = array1[pointer1]
				result[1] = array2[pointer2]
			}
			pointer1++
		} else {
			currentDifference := int(math.Abs(float64(array1[pointer1]) - float64(array2[pointer2])))
			if currentDifference < smallestDifference {
				smallestDifference = currentDifference
				result[0] = array1[pointer1]
				result[1] = array2[pointer2]
			}
			pointer2++
		}
	}

	return result
}

func SmallestDifferenceWithSortingAndTwoPointersImproved(array1, array2 []int) []int {
	smallestDifference := math.MaxInt64
	result := make([]int, 2)

	sort.Ints(array1)
	sort.Ints(array2)

	pointer1 := 0
	pointer2 := 0

	for pointer1 < len(array1) && pointer2 < len(array2) {
		currentDifference := int(math.Abs(float64(array1[pointer1]) - float64(array2[pointer2])))

		if currentDifference < smallestDifference {
			smallestDifference = currentDifference
			result[0] = array1[pointer1]
			result[1] = array2[pointer2]
		}

		if array1[pointer1] < array2[pointer2] {
			pointer1++
		} else {
			pointer2++
		}
	}

	return result
}
