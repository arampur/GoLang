package main

import "fmt"

func main() {
	inputMatrix := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
	}

	result := spiralMatrix(inputMatrix)
	fmt.Println(result)
}

func spiralMatrix(inputMatrix [][]int) []int {
	m, n := len(inputMatrix), len(inputMatrix[0])

	res := make([]int, 0, m*n)

	left, right, top, bottom := 0, n-1, 0, m-1

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, inputMatrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, inputMatrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, inputMatrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, inputMatrix[i][left])
			}
			left++
		}
	}

	return res
}
