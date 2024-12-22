package arrays

import "testing"

func BenchmarkNonConstructibleChange(b *testing.B) {
	coins := []int{5, 7, 1, 1, 2, 3, 22}
	for i := 0; i < b.N; i++ {
		NonConstructibleChange(coins)
	}
}
