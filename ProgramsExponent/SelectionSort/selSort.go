package main

import "fmt"

func main() {
	a := []int{3, 45, 6, 7, 31, 1}
	selSort(a)
	fmt.Println(a)
}

func selSort(a []int) []int {
	n := len(a)

	for i := 0; i < n-1; i++ {
		minIdx := i

		for j := i + 1; j < n; j++ {
			if a[j] < a[minIdx] {
				minIdx = j
			}
		}

		if minIdx != i {
			temp := a[i]
			a[i] = a[minIdx]
			a[minIdx] = temp
		}
	}

	return a
}
