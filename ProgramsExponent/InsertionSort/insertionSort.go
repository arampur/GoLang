package main

import "fmt"

func main() {
	a := []int{3, 45, 6, 7, 31, 1, 47, 43, 0, -1, 4}
	insertSort(a)
	fmt.Println(a)
}

func insertSort(a []int) []int {
	if len(a) == 1 {
		return a
	}

	n := len(a)

	for i := 1; i < n; i++ {
		key := a[i]

		j := i - 1
		for j >= 0 && a[j] > key {
			a[j+1] = a[j]
			j--
		}
		a[j+1] = key
	}

	return a
}
