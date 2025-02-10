package main

import "fmt"

type Pair struct {
	char  rune
	count int
}

func main() {
	res := adjacentStrings("aaabbbacd", 3)
	fmt.Println(res)
}

func adjacentStrings(s string, k int) string {
	stack := []Pair{}

	for _, char := range s {
		if len(stack) == 0 || stack[len(stack)-1].char != char {
			stack = append(stack, Pair{char, 1})
		} else {
			stack[len(stack)-1].count++
			if stack[len(stack)-1].count == k {
				stack = stack[:len(stack)-1]
			}
		}
	}

	res := ""

	for _, p := range stack {
		for i := 0; i < p.count; i++ {
			res += string(p.char)
		}
	}

	return res
}
