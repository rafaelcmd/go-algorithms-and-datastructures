package main

import insertionsort "github.com/rafaelcmd/go-algorithms-and-datastructures/sorting/insertion_sort"

func main() {
	var arr1 = []int{5, 2, 4, 6, 1, 3}

	insertionsort.InsertionSort(arr1)
}
