package arrays

// Time: O(n) | Space: O(n)
func SpiralTraverse(array [][]int) []int {
	var result []int
	top := 0
	bottom := len(array) - 1
	left := 0
	right := len(array[0]) - 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, array[top][i])
		}
		top++

		for j := top; j <= bottom; j++ {
			result = append(result, array[j][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, array[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for j := bottom; j >= top; j-- {
				result = append(result, array[j][left])
			}
			left++
		}
	}

	return result
}

// Time: O(n) | Space: O(n)
func SpiralTraverseImproved(array [][]int) []int {
	rows, cols := len(array), len(array[0])
	result := make([]int, 0, rows*cols) // Preallocate capacity

	top, bottom := 0, rows-1
	left, right := 0, cols-1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, array[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			result = append(result, array[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, array[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, array[i][left])
			}
			left++
		}
	}

	return result
}
