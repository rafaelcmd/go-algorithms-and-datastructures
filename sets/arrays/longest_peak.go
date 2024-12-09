package arrays

//Time complexity: O(n)
//Space complexity: O(n)
func LongestPeak(array []int) int {
	peaks := []int{}
	longestPeak := []int{}

	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i] > array[i+1] {
			peaks = append(peaks, i)
		}
	}

	for _, v := range peaks {
		currentLongestPeak := []int{}
		leftIndex := v - 1
		rightIndex := v + 1
		currentLongestPeak = append(currentLongestPeak, array[v])

		for leftIndex >= 0 && array[leftIndex] < array[leftIndex+1] {
			if array[leftIndex] < array[leftIndex+1] {
				currentLongestPeak = append(currentLongestPeak, array[leftIndex])
				leftIndex--
			}
		}

		for rightIndex < len(array) && array[rightIndex] < array[rightIndex-1] {
			if array[rightIndex] < array[rightIndex-1] {
				currentLongestPeak = append(currentLongestPeak, array[rightIndex])
				rightIndex++
			}
		}

		if len(currentLongestPeak) >= 3 && len(currentLongestPeak) > len(longestPeak) {
			longestPeak = currentLongestPeak
		}
	}

	return len(longestPeak)
}

//Time complexity: O(n)
//Space complexity: O(1)
func LongestPeakOptimized(array []int) int {
	longestPeakLength := 0

	for i := 1; i < len(array)-1; i++ {
		if array[i] > array[i-1] && array[i] > array[i+1] {
			leftIndex := i - 1
			for leftIndex >= 0 && array[leftIndex] < array[leftIndex+1] {
				leftIndex--
			}

			rightIndex := i + 1
			for rightIndex < len(array) && array[rightIndex] < array[rightIndex-1] {
				rightIndex++
			}

			peakLength := rightIndex - leftIndex - 1
			if peakLength > longestPeakLength {
				longestPeakLength = peakLength
			}
		}
	}

	return longestPeakLength
}

//Time complexity: O(n)
//Space complexity: O(1)
func LongestPeakWithDifferentApproach(array []int) int {
	longestPeakLength := 0
	i := 1

	for i < len(array)-1 {
		isPeak := array[i] > array[i-1] && array[i] > array[i+1]

		if !isPeak {
			i++
			continue
		}

		leftIndex := i - 2

		for leftIndex >= 0 && array[leftIndex] < array[leftIndex+1] {
			leftIndex--
		}

		rightIndex := i + 2

		for rightIndex < len(array) && array[rightIndex] < array[rightIndex-1] {
			rightIndex++
		}

		currentPeakLength := rightIndex - leftIndex - 1

		if currentPeakLength > longestPeakLength {
			longestPeakLength = currentPeakLength
		}

		i = rightIndex
	}

	return longestPeakLength
}
