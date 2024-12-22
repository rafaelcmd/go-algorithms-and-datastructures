package arrays

// Time complexity: O(n) | Space complexity: O(1)
func SubSequence(array []int, sequence []int) bool {
	arrIdx := 0
	seqIdx := 0

	for arrIdx < len(array) && seqIdx < len(sequence) {
		if array[arrIdx] == sequence[seqIdx] {
			seqIdx++
		}
		arrIdx++
	}

	return seqIdx == len(sequence)
}

// Time complexity: O(n) | Space complexity: O(1)
func IsValidSubsequence(array []int, sequence []int) bool {
	counter := 0
	var subSequence []int

	for _, v := range array {
		currentNum := sequence[counter]

		if currentNum == v {
			subSequence = append(subSequence, currentNum)
			counter++
		}

		if len(subSequence) == len(sequence) {
			return true
		}
	}

	return false
}
