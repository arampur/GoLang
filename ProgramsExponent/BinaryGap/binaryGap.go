package main

import (
	"fmt"
)

func main() {
	res := binaryGap(1041)
	fmt.Println(res)
}

func binaryGap(n int) int {
	maxGap := 0
	currentGap := 0
	foundOne := false

	for n > 0 {
		if n&1 == 1 {
			if foundOne && currentGap > maxGap {
				maxGap = currentGap
			}
			foundOne = true
			currentGap = 0
		} else if foundOne {
			currentGap++
		}
		n >>= 1
	}
	return maxGap
}
