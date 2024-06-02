package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{4, 3, 2, 5, 4}
	num := findRepeatedInArray(arr)
	if num != -1 {
		fmt.Println("Repeated element in the array:", num)
	} else {
		fmt.Println("There are no Repeated element in the array")
	}
}

func findRepeatedInArray(arr []int) int {
	sort.Ints(arr)
	fmt.Println("Arrays after calling sorted..")
	fmt.Println(arr)
	i := 0
	for i < len(arr)-1 {
		if arr[i] == arr[i+1] {
			return arr[i]
		} else {
			i++
		}
	}
	return -1
}
