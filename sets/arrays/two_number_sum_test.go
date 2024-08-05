package arrays

import (
	"testing"
)

// Benchmark function for TwoNumberSum
func BenchmarkTwoNumberSumUsingBruteForce(b *testing.B) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	for i := 0; i < b.N; i++ {
		TwoNumberSumUsingBruteForce(array, target)
	}
}

// Benchmark function for TwoNumberSumOptimized
func BenchmarkTwoNumberSumUsingHashing(b *testing.B) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	for i := 0; i < b.N; i++ {
		TwoNumberSumUsingHashing(array, target)
	}
}

func BenchmarkTwoNumberSumUsingSortingAndTwoPointers(b *testing.B) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	for i := 0; i < b.N; i++ {
		TwoNumberSumUsingSortingAndTwoPointers(array, target)
	}
}
