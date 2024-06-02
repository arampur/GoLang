package main

import "fmt"

func main() {
	arr := []int{5, 0, 0, 0}
	//fmt.Println(hasGoodSubarray(arr, 8))
	fmt.Println(hasGoodSubarray(arr, 3))
}

func hasGoodSubarray(arr []int, k int) bool {

	l := 0
	r := len(arr) - 1

	for l < r {
		if (arr[l]+arr[r])%k == 0 && (arr[l] != 0 && arr[r] != 0) {
			return true
		}
		if r == l+1 {
			l++
		}
		l++
	}

	return false
}
