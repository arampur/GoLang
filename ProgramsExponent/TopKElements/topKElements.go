package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 3, 3, 3}
	k := 2

	res := topKFrequent(nums, k)
	fmt.Println(res)

	nums = []int{7, 7}
	k = 1
	res = topKFrequent(nums, k)
	fmt.Println(res)

}

func topKFrequent(nums []int, k int) []int {
	res := []int{}

	hMap := make(map[int]int)

	for _, val := range nums {
		hMap[val]++
	}

	fmt.Println(hMap)
	for key, val := range hMap {
		if val >= k {
			res = append(res, key)
		}
	}
	return res
}
