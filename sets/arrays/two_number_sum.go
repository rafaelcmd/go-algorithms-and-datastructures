package arrays

// Time complexity: O(n^2) | Space complexity: O(1)
func TwoNumberSum(array []int, target int) []int {
	counter := 0

	if len(array) > 1 {
		for counter <= len(array)-1 {
			for i := 0; i < len(array)-1; i++ {
				currentNumber := array[counter]

				if currentNumber != array[i] && currentNumber+array[i] == target {
					return []int{currentNumber, array[i]}
				}
			}

			counter++
		}
	}

	return []int{}
}

// Time complexity: O(n) | Space complexity: O(n)
func OptimizedTwoNumberSum(array []int, target int) []int {
	results := make(map[int]bool)

	for _, num := range array {
		potentialMatch := target - num

		if results[num] {
			return []int{potentialMatch, num}
		}

		results[potentialMatch] = true

	}

	return []int{}
}
