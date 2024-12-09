package arrays

import "testing"

func BenchmarkLongestPeak(b *testing.B) {
	array := []int{1, 2, 3, 4, 5, 1}
	for i := 0; i < b.N; i++ {
		LongestPeak(array)
	}
}

func BenchmarkLongestPeakOptimized(b *testing.B) {
	array := []int{1, 2, 3, 4, 5, 1}
	for i := 0; i < b.N; i++ {
		LongestPeakOptimized(array)
	}
}

func BenchmarkLongestPeakWithDifferentApproach(b *testing.B) {
	array := []int{1, 2, 3, 4, 5, 1}
	for i := 0; i < b.N; i++ {
		LongestPeakWithDifferentApproach(array)
	}
}
