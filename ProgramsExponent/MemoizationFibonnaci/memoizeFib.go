package main

import "fmt"

var res = make(map[int]int)

func main() {
	fibNum := memoizeFibonacci(5)
	fmt.Println("res:", fibNum)
}

func memoizeFibonacci(n int) int {
	if val, key := res[n]; key {
		return val
	}

	if n == 0 || n == 1 {
		res[n] = n
		return n
	}

	output := memoizeFibonacci(n-1) + memoizeFibonacci(n-2)
	res[n] = output

	return output
}
