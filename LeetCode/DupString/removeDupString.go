package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "deeedbbcccbdaa"
	m := make(map[rune]int)
	res := identifyAdjacent(str, m, 3)
	fmt.Println("Result: ", res)
}

func identifyAdjacent(s string, m map[rune]int, k int) string {
	var b strings.Builder

	for _, c := range s {
		_, ok := m[c]
		if ok {
			m[c]++
		} else {
			m[c] = 1
		}
	}

	fmt.Println("Map contents: ", m)
	for _, c := range s {
		if m[c] != k {
			b.WriteRune(c)
		}
	}
	return b.String()
}
