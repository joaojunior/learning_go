package dsa

import (
	"errors"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](sorted_numbers []T, number T) (int, error) {
	left := 0
	right := len(sorted_numbers) - 1
	index := -1

	for ; left <= right ; {
		mid := left + (right - left) / 2
		if number <= sorted_numbers[mid] {
			if number == sorted_numbers[mid] {
				index = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if index == -1 || sorted_numbers[index] != number{
		return index, errors.New("Number is not in the array")
	}

	return index, nil
}
