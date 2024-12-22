package arrays

import "testing"

func BenchmarkTransposeMatrix1(b *testing.B) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		TransposeMatrix1(matrix)
	}
}

func BenchmarkTransposeMatrix2(b *testing.B) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		TransposeMatrix2(matrix)
	}
}

func BenchmarkTransposeMatrix3(b *testing.B) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < b.N; i++ {
		TransposeMatrix3(matrix)
	}
}
