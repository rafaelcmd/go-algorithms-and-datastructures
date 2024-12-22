package arrays

import (
	"testing"
)

// Benchmark function for TournamentWinner
func BenchmarkTournamentWinner(b *testing.B) {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	for i := 0; i < b.N; i++ {
		TournamentWinner(competitions, results)
	}
}

// Benchmark function for TournamentWinnerOptimized
func BenchmarkTournamentWinnerOptimized(b *testing.B) {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	for i := 0; i < b.N; i++ {
		TournamentWinnerOptimized(competitions, results)
	}
}

// Benchmark function for TournamentWinnerOptimizedImproved
func BenchmarkTournamentWinnerOptimizedImproved(b *testing.B) {
	competitions := [][]string{
		{"HTML", "C#"},
		{"C#", "Python"},
		{"Python", "HTML"},
	}
	results := []int{0, 0, 1}
	for i := 0; i < b.N; i++ {
		TournamentWinnerOptimizedImproved(competitions, results)
	}
}
