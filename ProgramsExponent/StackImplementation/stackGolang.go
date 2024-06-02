package main

import "fmt"

type Stack []int //LIFO structure

func main() {
	var stack Stack
	stack.Push(12)
	stack.Push(23)
	stack.Push(32)
	stack.Push(41)

	for len(stack) > 0 {
		x, y := stack.Pop()
		if y {
			fmt.Println("element popped:", x)
		}
	}
}

// Push item to stack
func (s *Stack) Push(element int) {
	*s = append(*s, element)
}

// Pop item from stack and return top of stack
func (s *Stack) Pop() (int, bool) {
	if s.isEmpty() {
		return -1, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

// Size of stack
func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}
