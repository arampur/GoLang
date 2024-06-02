package main

import "fmt"

func moveToEnd(arr []int) []int {
	if arr == nil || len(arr) < 2 {
		return arr
	}

	counter := 0
	for _, val := range arr {
		if val != 0 {
			arr[counter] = val
			counter++
		}
	}

	for counter < len(arr) {
		arr[counter] = 0
		counter++
	}

	return arr
}

func main() {
	example := []int{0, 1, 0, 3, 12}
	fmt.Println("Given array: ", example)
	moveToEnd(example)
	fmt.Println("After the shift of Zeros to end of array..")
	for _, num := range example {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}
