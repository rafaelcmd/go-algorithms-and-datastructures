package arrays

import "testing"

func BenchmarkThreeNumberSumWithSortingAndTwoPointers(b *testing.B) {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	for i := 0; i < b.N; i++ {
		ThreeNumberSumWithSortingAndTwoPointers(array, target)
	}
}

func BenchmarkThreeNumberSumWithSortingAndTwoPointersOptimized1(b *testing.B) {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	for i := 0; i < b.N; i++ {
		ThreeNumberSumWithSortingAndTwoPointersOptimized1(array, target)
	}
}

func BenchmarkThreeNumberSumWithSortingAndTwoPointersOptimized2(b *testing.B) {
	array := []int{12, 3, 1, 2, -6, 5, -8, 6}
	target := 0
	for i := 0; i < b.N; i++ {
		ThreeNumberSumWithSortingAndTwoPointersOptimized2(array, target)
	}
}
