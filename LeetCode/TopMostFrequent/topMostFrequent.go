//Given an integer array nums and an integer k, return the k most frequent elements.
//You may return the answer in any order.

// Example 1:

// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
// Example 2:

// Input: nums = [1], k = 1
// Output: [1]

package main

import (
	"container/heap"
	"fmt"
	"sort"
)

//type MaxHeap []int

type heapNode struct {
	num   int
	count int
}

type minHeap []heapNode

func (m minHeap) Len() int {
	return len(m)
}

func (m minHeap) Less(x, y int) bool {
	return m[x].count < m[y].count
}

func (m *minHeap) Push(x interface{}) {
	*m = append(*m, x.(heapNode))
}

func (m minHeap) Swap(x, y int) {
	m[x], m[y] = m[y], m[x]
}

func (m *minHeap) Pop() interface{} {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

// func (m MaxHeap) Len() int {
// 	return len(m)
// }

// func (m MaxHeap) Less(x, y int) bool {
// 	return m[x] > m[y]
// }

// func (m *MaxHeap) Push(x interface{}) {
// 	*m = append(*m, x.(int))
// }

// func (m MaxHeap) Swap(x, y int) {
// 	m[x], m[y] = m[y], m[x]
// }

// func (m *MaxHeap) Pop() interface{} {
// 	x := (*m)[len(*m)-1]
// 	*m = (*m)[:len(*m)-1]
// 	return x
// }

func topKFrequent(nums []int, k int) []int {
	// maxHeap := &MaxHeap{}
	// *maxHeap = nums
	// heap.Init(maxHeap)

	// res := []int{} // for top k most frequent

	// fmt.Println("MaxHeap:", maxHeap)

	// for i := 0; i < len(nums)-k; i++ {
	// 	heap.Pop(maxHeap)
	// }

	// fmt.Println("MaxHeap after:", maxHeap)

	// return res

	// put the key and values which includes frequency into map
	m := make(map[int]int)

	res := []int{}

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	//fmt.Println("map: ", m)

	keys := []int{} // to store the keys

	for key := range m {
		keys = append(keys, key)
	}

	//fmt.Println("keys : ", keys)

	// order the keys by descending order of frequency/count

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	//fmt.Println("After sorting keys : ", keys)

	// return first k elements

	for _, val := range keys {
		if len(res) == k {
			return res
		}
		res = append(res, val)
	}

	//return res

	//This program gives time complexity of O(N log N)
	// Can be improved using heap (min-heap)

	// Using min-heap that calculates sorts based on the least frequency
	// Reduces the time complexity => O(N log K)

	h := &minHeap{}
	heap.Init(h)

	for num, count := range m {
		heap.Push(h, heapNode{num: num, count: count})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	newRes := make([]int, 0, k)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(heapNode).num)
	}
	return newRes
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	res := topKFrequent(nums, 2)

	fmt.Println("top k: ", res)
}
