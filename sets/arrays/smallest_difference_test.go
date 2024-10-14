package arrays

import "testing"

func BenchmarkSmallestDifference(b *testing.B) {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for i := 0; i < b.N; i++ {
		SmallestDifference(array1, array2)
	}
}

func BenchmarkSmallestDifferenceWithSortingAndTwoPointers(b *testing.B) {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for i := 0; i < b.N; i++ {
		SmallestDifferenceWithSortingAndTwoPointers(array1, array2)
	}
}

func BenchmarkSmallestDifferenceWithSortingAndTwoPointersImproved(b *testing.B) {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array2 := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	for i := 0; i < b.N; i++ {
		SmallestDifferenceWithSortingAndTwoPointersImproved(array1, array2)
	}
}
