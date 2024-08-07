package arrays

import (
	"testing"
)

// Benchmark function for TwoNumberSum
func BenchmarkSubSequence(b *testing.B) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	for i := 0; i < b.N; i++ {
		SubSequence(array, sequence)
	}
}

// Benchmark function for TwoNumberSumOptimized
func BenchmarkIsValidSubsequence(b *testing.B) {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	for i := 0; i < b.N; i++ {
		IsValidSubsequence(array, sequence)
	}
}
