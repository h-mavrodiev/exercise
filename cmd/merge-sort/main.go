package main

import "fmt"

func MergeSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}

	halfLen := len(unsorted) / 2

	a := MergeSort(unsorted[halfLen:])
	b := MergeSort(unsorted[:halfLen])

	return merge(a, b)
}

func merge(a, b []int) []int {
	var (
		merged []int
		i, j   int
	)

	for len(a) > i && len(b) > j {
		if a[i] > b[j] {
			merged = append(merged, a[i])
			i++
		} else if a[i] < b[j] {
			merged = append(merged, b[j])
			j++
		}
	}

	// take care of leftovers in both slices
	for ; len(a) > i; i++ {
		merged = append(merged, a[i])
	}

	for ; len(b) > j; j++ {
		merged = append(merged, b[j])
	}
	return merged
}

func main() {
	var unsorted []int = []int{10, 6, 2, 1, 5, 8, 3, 4, 7, 9} // []int{3, 6, 1, 8, 9, 12, 2, 7}
	sorted := MergeSort(unsorted)
	fmt.Printf("unsorted: %v\n", unsorted)
	fmt.Printf("sorted: %v\n", sorted)
}
