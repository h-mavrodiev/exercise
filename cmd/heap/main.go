package main

import (
	"fmt"
	"math/rand"
)

type MaxHeap struct {
	array []int
}

func NewHeap() *MaxHeap {
	return &MaxHeap{}
}

func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

func (h *MaxHeap) Extract() int {
	lastIdx := len(h.array) - 1
	r := h.array[0]
	h.array[0] = h.array[lastIdx]
	h.array = h.array[:lastIdx]

	h.maxHeapifyDown(0)

	return r
}

func (h *MaxHeap) maxHeapifyUp(i int) {
	for h.array[parent(i)] < h.array[i] {
		h.swap(i, parent(i))
		i = parent(i) //update the loop
	}
}

func (h *MaxHeap) maxHeapifyDown(i int) {
	var (
		toSwapIdx int
		lastIdx   int = len(h.array) - 1
		li, ri    int = left(i), right(i)
	)

	for li <= lastIdx {
		if li == lastIdx {
			toSwapIdx = li
		} else if h.array[li] > h.array[i] {
			toSwapIdx = li
		} else {
			toSwapIdx = ri
		}

		if h.array[toSwapIdx] > h.array[i] {
			h.swap(toSwapIdx, i)
			i = toSwapIdx
			li, ri = left(i), right(i)
		} else {
			return
		}
	}

}

func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func main() {
	h := &MaxHeap{}

	for _, v := range rand.Perm(12) {
		h.Insert(v)
		fmt.Println(h.array)
	}

	for i := 0; i < 5; i++ {
		h.Extract()
		fmt.Println(h.array)
	}
}
