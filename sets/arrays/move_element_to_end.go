package arrays

// Time complexity: O(n)
// Space complexity: O(1)
func MoveElementToEndWithTwoPointers(array []int, toMove int) []int {
	first := 0
	last := len(array) - 1

	for first < last {
		if array[first] == toMove && array[last] != toMove {
			array[first] = array[last]
			array[last] = toMove
			last--
			first++
		} else if array[first] == toMove && array[last] == toMove {
			last--
		} else if array[first] != toMove && array[last] == toMove {
			last--
			first++
		} else if array[first] != toMove && array[last] != toMove {
			first++
		}
	}

	return array
}

// Time complexity: O(n)
// Space complexity: O(1)
func MoveElementToEndWithTwoPointersImproved(array []int, toMove int) []int {
	first := 0
	last := len(array) - 1

	for first < last {
		if array[last] == toMove {
			last--
		} else if array[first] == toMove {
			array[first], array[last] = array[last], array[first]
			first++
			last--
		} else {
			first++
		}
	}

	return array
}

// Time complexity: O(n)
// Space complexity: O(1)
func MoveElementToEndWithTwoPointersImproved2(array []int, toMove int) []int {
	first := 0
	last := len(array) - 1

	for first < last {
		for first < last && array[last] == toMove {
			last--
		}

		if array[first] == toMove {
			array[first], array[last] = array[last], array[first]
		}

		first++
	}

	return array
}
