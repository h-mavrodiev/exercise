package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, startIdx, endIdx int) []int {
	if startIdx < endIdx {
		var pivotIdx int

		arr, pivotIdx = partition(arr, startIdx, endIdx)
		arr = quickSort(arr, startIdx, pivotIdx-1)
		arr = quickSort(arr, pivotIdx+1, endIdx)
	}

	return arr
}

func partition(arr []int, startIdx, endIdx int) ([]int, int) {
	toCompare := arr[endIdx]
	newPivotIdx := startIdx

	for i := newPivotIdx; i < endIdx; i++ {
		if arr[i] < toCompare {
			arr[i], arr[newPivotIdx] = arr[newPivotIdx], arr[i]
			newPivotIdx++
		}
	}

	arr[newPivotIdx], arr[endIdx] = arr[endIdx], arr[newPivotIdx]
	return arr, newPivotIdx
}

func main() {
	rand := rand.Perm(12)

	sort := QuickSort(rand)
	fmt.Println(sort)
}
