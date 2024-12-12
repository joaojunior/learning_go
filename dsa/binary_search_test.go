package dsa

import "testing"

func TestBinarySearchInAnEmptyArray(t *testing.T){
	empty := []int{}
	_, err := BinarySearch(empty, 42)

	if err == nil {
		t.Fatalf("BinarySearch is not working in an empty array")
	}
}

func TestBinarySearchInArraySizeOneWithNumberIsInTheArray(t *testing.T){
	number := 42
	numbers := []int{number}
	index, _ := BinarySearch(numbers, number)

	if index != 0 {
		t.Fatalf("BinarySearch does not found the correct index: expected=%v, found=%v", 0, index)
	}
}

func TestBinarySearchInArraySizeOneWithNumberIsNotInTheArray(t *testing.T){
	number := 42
	numbers := []int{number}
	index, _ := BinarySearch(numbers, -number)

	if index != -1 {
		t.Fatalf("BinarySearch does not found the correct index: expected=%v, found=%v", -1, index)
	}
}

func TestBinarySearchInArraySizeEvenWithNumberIsInTheArray(t *testing.T){
	number := 42
	numbers := []int{number, number + 1, number + 10, 100}
	index, _ := BinarySearch(numbers, number)

	if index != 0 {
		t.Fatalf("BinarySearch does not found the correct index: expected=%v, found=%v", 0, index)
	}
}

func TestBinarySearchInArraySizeOddWithNumberIsInTheArray(t *testing.T){
	number := 42
	numbers := []int{number, number + 1, number + 10}
	index, _ := BinarySearch(numbers, number)

	if index != 0 {
		t.Fatalf("BinarySearch does not found the correct index: expected=%v, found=%v", 0, index)
	}
}

func TestBinarySearchInArraySizeEvenWithNumberIsNotInTheArray(t *testing.T){
	number := 42
	numbers := []int{number - 1, number + 1, number + 10, 100}
	index, err := BinarySearch(numbers, number)

	if err == nil {
		t.Fatalf("BinarySearch does not found the correct index: expected=%v, found=%v", 0, index)
	}
}

func TestBinarySearchInArraySizeOddWithNumberIsNotInTheArray(t *testing.T){
	number := 42
	numbers := []int{number - 1, number + 1, number + 10}
	index, err := BinarySearch(numbers, number)

	if err == nil {
		t.Fatalf("BinarySearch does not found the correct index: expected=%v, found=%v", -1, index)
	}
}
