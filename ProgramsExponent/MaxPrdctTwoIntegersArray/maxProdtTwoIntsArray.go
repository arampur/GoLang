package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-152, 22, -3}
	result := findMaxProductTwoIntArrays(arr)
	fmt.Println("Final product ", result)
}

func findMaxProductTwoIntArrays(arr []int) int {

	n := len(arr)

	if n == 1 {
		return -1
	} else if n == 2 {
		return arr[0] * arr[1]
	}

	sort.Ints(arr)

	if arr[n-1]*arr[n-2] > arr[0]*arr[1] {
		return arr[n-1] * arr[n-2]
	} else {
		return arr[0] * arr[1]
	}
}
