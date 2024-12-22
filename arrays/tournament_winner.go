package arrays

// Time complexity: O(n + m) | Space complexity: O(m)
func TournamentWinner(competitions [][]string, results []int) string {
	result := map[string]int{}

	for k, v := range results {
		if v == 0 {
			result[competitions[k][1]] = result[competitions[k][1]] + 3
		} else {
			result[competitions[k][0]] = result[competitions[k][0]] + 3
		}
	}

	return winner(result)
}

func winner(winners map[string]int) string {
	winner := ""
	higherPoints := -1

	for k := range winners {
		if winners[k] > higherPoints {
			higherPoints = winners[k]
			winner = k
		}
	}

	return winner
}

// Time complexity: O(n) | Space complexity: O(m)
func TournamentWinnerOptimized(competitions [][]string, results []int) string {
	result := map[string]int{}
	winner := ""
	higherPoints := -1

	for k, v := range results {
		if v == 0 {
			currentWinner := competitions[k][1]
			result[currentWinner] = result[currentWinner] + 3
			if result[currentWinner]+3 > higherPoints {
				higherPoints = result[currentWinner]
				winner = currentWinner
			}
		} else {
			currentWinner := competitions[k][0]
			result[currentWinner] = result[currentWinner] + 3
			if result[currentWinner]+3 > higherPoints {
				higherPoints = result[currentWinner]
				winner = currentWinner
			}
		}
	}

	return winner
}

// Time complexity: O(n) | Space complexity: O(m)
func TournamentWinnerOptimizedImproved(competitions [][]string, results []int) string {
	scores := map[string]int{}
	winner := ""
	highestPoints := 0

	for i, result := range results {
		winningTeam := competitions[i][1]
		if result == 1 {
			winningTeam = competitions[i][0]
		}

		scores[winningTeam] += 3

		if scores[winningTeam] > highestPoints {
			highestPoints = scores[winningTeam]
			winner = winningTeam
		}
	}

	return winner
}
