package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 19, 5, -4, 7, 18, 15, -10}
	res := maxSubArray(arr)
	fmt.Println(res)

}

func maxSubArray(arr []int) int {
	n := len(arr)
	maxDiff := 0
	for i := 1; i < n; i++ {
		for j := i + 1; j < n; j++ {
			maxDiff = max(maxDiff, arr[j]-arr[i])
		}
	}
	return maxDiff
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// func min(a int, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func maxSubArray2(arr []int, l int, u int) int {
// 	if u == l {
// 		return 0
// 	} else if u == l+1 {
// 		return max(arr[u]-arr[l], 0)
// 	}

// 	m := (l + u) / 2

// 	m1 := maxSubArray2(arr, l, m)
// 	m2 := maxSubArray2(arr, m+1, u)

// 	y1 := maxElement(arr, m+1, u)
// 	x1 := minElement(arr, l, m)

// 	firstmax := max(m1, m2)
// 	return max(firstmax, y1-x1)
// }

// func minElement(arr []int, l int, u int) int {
// 	minElement := math.MaxInt32

// 	for i := l; i < u; i++ {
// 		minElement = min(minElement, arr[i])
// 	}
// 	return minElement
// }

// func maxElement(arr []int, l int, u int) int {
// 	maxElement := math.MinInt32

// 	for i := l; i < u; i++ {
// 		maxElement = max(maxElement, arr[i])
// 	}
// 	return maxElement
// }
