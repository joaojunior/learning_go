package dsa

import (
	"testing"
	"reflect"
	"slices"
)

func TestMergesortEmptyArray(t *testing.T){
	empty := []int{}
	sorted := MergeSort(empty)

	if len(sorted) != 0 {
		t.Fatalf("Mergesort does not sort an empty array: %v", sorted)
	}
}

func TestMergesortSortsArrayWithOneElement(t *testing.T){
	data := []int{1}
	sorted := MergeSort(data)

	if !reflect.DeepEqual(data, sorted) {
		t.Fatalf("Mergesort does not sort an array with one element: input=%v,output=%v", data, sorted)
	}
}

func TestMergesortSortsArrayAlreadySorted(t *testing.T){
	data := []int{1,2,3,4,5}

	sorted := MergeSort(data)

	if !reflect.DeepEqual(data, sorted) {
		t.Fatalf("Mergesort does not sort an array already sorted: input=%v,output=%v", data, sorted)
	}
}

func TestMergesortSortsArrayAlreadySortedDesc(t *testing.T){
	data := []int{5,4,3,2,1}

	sorted := MergeSort(data)
	slices.Sort(data)

	if !reflect.DeepEqual(data, sorted) {
		t.Fatalf("Mergesort does not sort an array already sorted desc: input=%v,output=%v", data, sorted)
	}
}

func TestMergesortSortsAnRandomArray(t *testing.T){
	data := []int{2,1,5,3,4}

	sorted := MergeSort(data)
	slices.Sort(data)

	if !reflect.DeepEqual(data, sorted) {
		t.Fatalf("Mergesort does not sort an random array: input=%v,output=%v", data, sorted)
	}
}
