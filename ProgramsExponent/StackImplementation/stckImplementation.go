package main

import "fmt"

type stack []int

func (s *stack) pop() (int, bool) {
	if s.isEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1
		item := (*s)[index]
		*s = (*s)[:index]
		return item, true
	}
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(item int) {
	*s = append(*s, item)
}

func main() {
	var stck stack
	stck.push(1)
	stck.push(2)
	stck.push(5)
	stck.push(6)

	stck.pop()
	stck.pop()

	fmt.Println(stck)

}
