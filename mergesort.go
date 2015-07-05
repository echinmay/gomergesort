package main

import (
	"fmt"
)

func merge(left, right []int) []int {
	retarr := make([]int, 0)

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			retarr = append(retarr, left[0])
			left = left[1:]
		} else {
			retarr = append(retarr, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		retarr = append(retarr, left...)
	}

	if len(right) > 0 {
		retarr = append(retarr, right...)
	}

	return retarr
}

func mergesort(arr []int) []int {
	lenofarr := len(arr)

	if lenofarr < 2 {
		return arr
	}

	mid := lenofarr / 2

	return merge(mergesort(arr[:mid]), mergesort(arr[mid:]))
}

func main() {
	arr := make([]int, 0)

	arr = append(arr, 37, 7, 2, 14, 35, 47, 10, 24, 44, 17, 34, 11, 16, 48, 1, 39, 6, 33, 43, 26, 40, 4, 28, 5, 38, 41, 42, 12, 13, 21, 29, 18, 3, 19, 0, 32, 46, 27, 31, 25, 15, 36, 20, 8, 9, 49, 22, 23, 30, 45)
	fmt.Println("Input ", arr)
	merged := mergesort(arr)
	fmt.Println(merged)
}
