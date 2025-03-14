package main

import "fmt"

// This program is to find the previous larger element for a given element
// if its first element then there is no larger element, therefore first element in result is always -1
// for any other element if no larger prev element exists, then the answer will be -1 if not, it will
// be next largeest element on the left

func main() {
	arr := []int{9, 6, 10, 9, 5}
	prevGre := prevGreatestElement(arr)
	fmt.Println("prevsmall greatest elements:", prevGre)
}

func prevGreatestElement(arr []int) []int {
	stack := []int{}
	res := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] <= arr[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
		stack = append(stack, arr[i])
	}
	return res
}
