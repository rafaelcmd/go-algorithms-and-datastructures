package arrays

// Time: O(n * m) | Space: O(n * m)
func TransposeMatrix1(matrix [][]int) [][]int {
	counter := 0
	var transposedMatrix [][]int

	if len(matrix) == len(matrix[0]) {
		transposedMatrix = make([][]int, len(matrix))

		for i := range transposedMatrix {
			transposedMatrix[i] = make([]int, len(matrix[0]))
		}

		for i := 0; i < len(matrix); i++ {
			for counter < len(matrix[0]) {
				transposedMatrix[i][counter] = matrix[counter][i]
				counter++
			}
			counter = 0
		}
	} else {
		transposedMatrix = make([][]int, len(matrix[0]))

		for i := range transposedMatrix {
			transposedMatrix[i] = make([]int, len(matrix))
		}

		for i := 0; i < len(matrix[0]); i++ {
			for counter < len(matrix) {
				transposedMatrix[i][counter] = matrix[counter][i]
				counter++
			}

			counter = 0
		}
	}

	return transposedMatrix
}

// Time: O(n * m) | Space: O(n * m)
func TransposeMatrix2(matrix [][]int) [][]int {
	transposedMatrix := make([][]int, len(matrix[0]))
	counter := 0

	for i := range transposedMatrix {
		transposedMatrix[i] = make([]int, len(matrix))

		var newRow []int

		for j := range matrix {
			newRow = append(newRow, matrix[j][i])

			if len(newRow) == len(matrix) {
				transposedMatrix[counter] = newRow
				counter++
			}
		}
	}

	return transposedMatrix
}

// Time: O(n * m) | Space: O(n * m)
func TransposeMatrix3(matrix [][]int) [][]int {
	transposedMatrix := make([][]int, len(matrix[0]))

	for i := range transposedMatrix {
		transposedMatrix[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			transposedMatrix[j][i] = matrix[i][j]
		}
	}

	return transposedMatrix
}
