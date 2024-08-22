package arrays

import (
	"testing"
)

// Benchmark function for SortedSquaredArray
func BenchmarkSortedSquaredArray(b *testing.B) {
	array := []int{-10, -5, 0, 5, 10}
	for i := 0; i < b.N; i++ {
		SortedSquaredArray(array)
	}
}

// Benchmark function for SortedSquaredArrayTwoPointers
func BenchmarkSortedSquaredArrayTwoPointers(b *testing.B) {
	array := []int{-10, -5, 0, 5, 10}
	for i := 0; i < b.N; i++ {
		SortedSquaredArrayTwoPointers(array)
	}
}

// Benchmark function for SortedSquaredArrayTwoPointersOptimized
func BenchmarkSortedSquaredArrayTwoPointersOptimized(b *testing.B) {
	array := []int{-10, -5, 0, 5, 10}
	for i := 0; i < b.N; i++ {
		SortedSquaredArrayTwoPointersOptimized(array)
	}
}
