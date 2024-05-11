package insertionvsmergesort

import (
	"math/rand"
	"time"
)

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge1(left, right)
}

func merge1(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}

func generateRandomArray(size int) []int {
	arr := make([]int, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(1000)
	}
	return arr
}

/*func main() {
	sizes := []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000}
	insertionTimes := make([]float64, len(sizes))
	mergeTimes := make([]float64, len(sizes))

	for i, size := range sizes {
		arr := generateRandomArray(size)

		// Measure time taken for insertion sort
		start := time.Now()
		insertionSort(arr)
		insertionTimes[i] = time.Since(start).Seconds()

		// Measure time taken for merge sort
		start = time.Now()
		sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
		mergeTimes[i] = time.Since(start).Seconds()
	}

	// Print results
	fmt.Println("Input Size\tInsertion Sort\tMerge Sort")
	for i, size := range sizes {
		fmt.Printf("%d\t\t%.6f\t\t%.6f\n", size, insertionTimes[i], mergeTimes[i])
	}
}*/
