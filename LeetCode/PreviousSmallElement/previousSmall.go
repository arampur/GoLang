package main

import "fmt"

// This program is to find the previous small element for a given element
// if its first element then there is no smaller element, therefore first element in result is always -1
// for any other element if no smaller prev element exists, then the answer will be -1 if not, it will
// be next smallest element on the left
func main() {
	arr := []int{9, 6, 10, 9, 5}
	prevSmallE := prevSmall(arr)
	fmt.Println("prevsmall elements:", prevSmallE)
}

func prevSmall(arr []int) []int {
	res := make([]int, len(arr))
	stack := []int{}

	//	arr := []int{9, 6, 10, 9, 5}

	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && stack[len(stack)-1] >= arr[i] {
			stack = stack[:len(stack)-1] // Pop the top element
		}

		// If stack is empty, there is no smaller element
		if len(stack) == 0 {
			res[i] = -1
		} else { // you have smaller element and add it to the result
			res[i] = stack[len(stack)-1]
		}

		// Push current element to stack
		stack = append(stack, arr[i])
	}
	return res
}
