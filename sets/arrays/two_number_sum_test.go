package arrays

import (
	"testing"
)

// Benchmark function for TwoNumberSum
func BenchmarkTwoNumberSum(b *testing.B) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	for i := 0; i < b.N; i++ {
		TwoNumberSum(array, target)
	}
}

// Benchmark function for TwoNumberSumOptimized
func BenchmarkTwoNumberSumOptimized(b *testing.B) {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	target := 10
	for i := 0; i < b.N; i++ {
		OptimizedTwoNumberSum(array, target)
	}
}
