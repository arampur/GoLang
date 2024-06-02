package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 2, 3, 7, 8, 6, 10, 9}
	fmt.Println(SortKMessedArray(arr,2))
}

func SortKMessedArray(arr []int, k int) []int {
	// your code goes here
	for i:=1;i<len(arr);i++ {
		x := arr[i]
		j := i-1

		for j >= 0 && arr[j] > x {
			arr[j+1] = arr[j]
			j--
		}

		moves := 0

		if moves >= k {
			break
		} else {
			moves++
		}

		arr[j+1] = x
	}

	return arr
}
