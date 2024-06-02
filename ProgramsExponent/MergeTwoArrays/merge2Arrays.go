package main

import "fmt"

func main() {
	a1 := []int{2, 3, 4, 5, 6, 7}
	a2 := []int{4, 5, 6, 7, 8, 9}

	res := mergeTwoArrays(a1, a2)
	fmt.Println(res)
}

func mergeTwoArrays(a1 []int, a2 []int) []int {
	res := []int{}

	for i := 0; i < len(a1); i++ {
		res = append(res, a1[i])
	}
	for i := 0; i < len(a2); i++ {
		res = append(res, a2[i])
	}

	return res
}
