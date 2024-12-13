package dsa

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](data []T) []T {
	if len(data) > 1 {
		mid := len(data) / 2
		left := MergeSort(data[0:mid])
		right := MergeSort(data[mid:])
		return merge(left, right)
	}

	return data
}

func merge[T constraints.Ordered](left []T, right []T) []T {
	var result []T
	left_index := 0
	right_index := 0

	for ; right_index < len(right) && left_index < len(left) ; {
		if left[left_index] <= right[right_index] {
			result = append(result, left[left_index])
			left_index += 1
		} else {
			result = append(result, right[right_index])
			right_index += 1
		}
	}

	for ; left_index < len(left) ; left_index++ {
		result = append(result, left[left_index])
	}

	for ; right_index < len(right) ; right_index++ {
		result = append(result, right[right_index])
	}

	return result
}
