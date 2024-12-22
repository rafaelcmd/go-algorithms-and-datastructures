package arrays

import "testing"

func BenchmarkMoveElementToEndWithTwoPointers(b *testing.B) {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	for i := 0; i < b.N; i++ {
		MoveElementToEndWithTwoPointers(array, toMove)
	}
}

func BenchmarkMoveElementToEndWithTwoPointersImproved(b *testing.B) {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	for i := 0; i < b.N; i++ {
		MoveElementToEndWithTwoPointersImproved(array, toMove)
	}
}

func BenchmarkMoveElementToEndWithTwoPointersImproved2(b *testing.B) {
	array := []int{2, 1, 2, 2, 2, 3, 4, 2}
	toMove := 2
	for i := 0; i < b.N; i++ {
		MoveElementToEndWithTwoPointersImproved2(array, toMove)
	}
}
