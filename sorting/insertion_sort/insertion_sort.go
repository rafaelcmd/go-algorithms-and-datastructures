package insertionsort

import "fmt"

func InsertionSort(arr []int) {
	fmt.Println(arr)

	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1

		for j > -1 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}

	fmt.Println(arr)
}
