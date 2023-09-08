package main

import (
	"fmt"
	"math/rand"
)

func InsertSort(arr []int) []int {
	// -------> moves to the right
	for i := 1; i < len(arr); i++ {
		// <------- moves to the left
		for j := i; j > 0 && arr[j-1] < arr[j]; j-- {
			arr[j], arr[j-1] = arr[j-1], arr[j]
		}
	}
	return arr
}

func BubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] < arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}

func main() {
	unsorted1 := rand.Perm(10)
	fmt.Println(unsorted1)
	sorted1 := InsertSort(unsorted1)
	fmt.Println(sorted1)

	unsorted2 := rand.Perm(10)
	fmt.Println(unsorted2)
	sorted2 := BubbleSort(unsorted2)
	fmt.Println(sorted2)
}
