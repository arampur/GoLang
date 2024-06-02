package main

import "fmt"

func main() {
	a := []int{0, 4, 0, 5, 10, 0, 3, 0}
	result := moveZerosAtEnd(a)
	fmt.Println(result)
}

func moveZerosAtEnd(arr []int) []int {
	n := len(arr)

	if n == 1 {
		return arr
	}

	l := 0
	r := n - 1

	for l < r {
		if arr[l] == 0 && arr[r] == 0 {
			for arr[r] == 0 && l < r {
				r--
			}

			if l < r {
				temp := arr[r]
				arr[r] = arr[l]
				arr[l] = temp
				l++
			}
		}
		if arr[l] != 0 && arr[r] == 0 {
			r--
			l++
		}

		if arr[l] == 0 && arr[r] != 0 {
			temp := arr[r]
			arr[r] = arr[l]
			arr[l] = temp
			l++
			r--
		}

		if arr[l] != 0 && arr[r] != 0 {
			l++
			r--
		}
	}

	return arr

}
