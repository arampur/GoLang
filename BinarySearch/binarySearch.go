package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var arr1 = []int{1, 3, 4, 5, 6, 7, 8, 10}
	key := 8
	res := doBinaryIndexSearch(arr1, key)
	fmt.Println("result: ", res)

	fmt.Println("Validating IP problem")
	s1 := validateIP("123.24.59.99")
	fmt.Println(s1)
}

func doBinaryIndexSearch(arr []int, x int) int {
	low := 0
	high := len(arr) - 1

	if arr[0] >= x {
		return 0
	}

	if len(arr)-1 < x {
		return -1
	}

	for low < high {
		mid := (low + high) / 2
		if arr[mid] <= x {
			low = mid
		} else {
			high = mid
		}
	}
	return high
}

func validateIP(s string) bool {
	str := strings.Split(s, ".")
	fmt.Println(str)

	if len(str) != 4 {
		return false
	}

	for i := 0; i < len(str); i++ {
		s1, err := strconv.Atoi(str[i])
		if err == nil {
			if s1 < 0 || s1 > 255 {
				return false
			}
		}
	}
	return true
}
