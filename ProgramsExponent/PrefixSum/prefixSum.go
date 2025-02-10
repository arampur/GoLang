package main

import "fmt"

func main() {
	// A -> 1 2 3 4 5 6
	// Res -> 1 3 6 10 15 21

	a := []int{1, 2, 3, 4}
	res := prefSum(a)
	fmt.Println(res)
}

func prefSum(arr []int) []int {
	res := []int{}

	// if len(arr) == 1 {
	// 	res = append(res, arr[0])
	// 	return res
	// }

	// sum := 0
	// for i := 1; i < n-1; i++ {
	// 	j := i
	// 	for j != 0 {
	// 		sum += arr[j]
	// 		fmt.Println("Sum: ", sum)
	// 		j--
	// 	}
	// 	res = append(res, sum)
	// 	sum = 0
	// }
	// return res
	sum := 0

	for i := 0; i < len(arr)-1; i++ {
		if i == 0 {
			res = append(res, arr[i])
		} else {
			j := i + 1
			for j != 0 {
				sum += arr[j]
				j--
			}
			res = append(res, sum)
			sum = 0
		}
	}

	return res
}
