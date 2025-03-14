//215 Program: Kth largest element in an array - Minheap

// Given an integer array nums and an integer k, return the kth largest element in the array.
// Note that it is the kth largest element in the sorted order, not the kth distinct element.
// Can you solve it without sorting?

package main

import (
	"container/heap"
	"fmt"
)

type minHeap []int

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(x, y int) bool {
	return h[x] < h[y]
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// func (h *minHeap) Pop() interface {} {
//     old := *h
//     n := len(old)
//     x := old[n-1]
//     *h = old[:n-1]
//     return x
// }

func (h *minHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	minimHeap := &minHeap{}
	*minimHeap = nums
	heap.Init(minimHeap)

	fmt.Println("minimHeap", minimHeap)

	for i := 0; i < len(nums)-k; i++ {
		heap.Pop(minimHeap)
	}

	fmt.Println("MinHeap after removal: ", minimHeap)

	return heap.Pop(minimHeap).(int)
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	fmt.Println("Given array:", nums)
	fmt.Println("MinHeap formed: ")
	x := findKthLargest(nums, 4)
	fmt.Println()
	fmt.Println("kth largest element: ", x)
}
