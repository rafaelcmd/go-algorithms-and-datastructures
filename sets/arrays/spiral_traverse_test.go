package arrays

import "testing"

func BenchmarkSpiralTraverse(b *testing.B) {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	for i := 0; i < b.N; i++ {
		SpiralTraverse(array)
	}
}

func BenchmarkSpiralTraverseImproved(b *testing.B) {
	array := [][]int{
		{1, 2, 3, 4},
		{12, 13, 14, 5},
		{11, 16, 15, 6},
		{10, 9, 8, 7},
	}

	for i := 0; i < b.N; i++ {
		SpiralTraverseImproved(array)
	}
}
