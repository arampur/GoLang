package main

import "fmt"

func main() {
	var arr1 = []int{1, 3, 4, 5, 6, 7, 8, 10}
	key := 9
	res := doBinaryIndexSearch(arr1, key)
	fmt.Println("result: ", res)
}

func doBinaryIndexSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	if arr[0] >= x {
		return 0
	}

	if len(arr)-1 < x {
		return -1
	}

	for low < high {
		mid := (low + high) / 2
		if arr[mid] <= x {
			low = mid
		} else {
			high = mid
		}
	}
	return high
}
