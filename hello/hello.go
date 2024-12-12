package main

import (
    "fmt"

    "example.com/dsa"
)

func main() {
	data := []int{1,4,0}
	message := dsa.Merge_sort(data[:])
	fmt.Println(data)
	fmt.Println(message)
}
