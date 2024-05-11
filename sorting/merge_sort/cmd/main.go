package main

import (
	"fmt"

	mergesort "github.com/rafaelcmd/go-algorithms-and-datastructures/sorting/merge_sort"
)

func main() {
	var arr = []int{2, 4, 1, 5, 3, 7, 8, 6}
	p := 0
	r := len(arr) - 1
	q := (p + r) / 2

	mergesort.Merge(arr, p, q, r)

	fmt.Println(arr)
}
