package main

import "fmt"

type Stack []string

func ReverseWords(arr []string) []string {
	// your code goes here
	var stck Stack
	res := []string{}

	for i := len(arr) - 1; i >= 0; i-- {
		stck.Push(arr[i])
	}

	fmt.Println("Stack:", stck)

	for len(stck) != 0 {
		val, _ := stck.Pop()
		res = append(res, val)
	}

	return res
}

func main() {
	arr := []string{"p", "e", "r", "f", "e", "c", "t", " ",
		"m", "a", "k", "e", "s", "  ",
		"p", "r", "a", "c", "t", "i", "c", "e"}

	//ecitcarp sekam tcefrep

	res := ReverseWords(arr)

	
	fmt.Println(res)

}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// func main() {
// 	var stack Stack // create a stack variable of type Stack

// 	stack.Push("this")
// 	stack.Push("is")
// 	stack.Push("sparta!!")

// 	for len(stack) > 0 {
// 		x, y := stack.Pop()
// 		if y == true {
// 			fmt.Println(x)
// 		}
// 	}
// }
