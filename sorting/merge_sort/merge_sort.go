package mergesort

func Merge(arr []int, p, q, r int) {
	lenLeft := q - p + 1
	lenRight := r - q

	left := make([]int, lenLeft)
	right := make([]int, lenRight)

	for i := 0; i < lenLeft; i++ {
		left[i] = arr[p+i]
	}

	for j := 0; j < lenRight; j++ {
		right[j] = arr[q+1+j]
	}

	i := 0
	j := 0
	k := p

	for i < lenLeft && j < lenRight {
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}
		k++
	}

	for i < lenLeft {
		arr[k] = left[i]
		i++
		k++
	}

	for j < lenRight {
		arr[k] = right[j]
		j++
		k++
	}
}
