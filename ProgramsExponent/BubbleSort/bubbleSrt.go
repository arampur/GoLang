package main

import "fmt"

func main() {
	a := []int{6, 5, 3, 1, 8, 7, 2, 4}
	bubbleSrt(a)
	fmt.Println(a)
}

func bubbleSrt(a []int) []int {
	n := len(a)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
			}
		}
	}

	return a
}
