package main

import "fmt"

type Stack []rune

func isBalanced(s string) bool {
	// Write your code here
	//map to store the bracket pairs
	m := make(map[rune]rune)
	var st Stack

	n := len(s)
	if n < 2 || n%2 != 0 {
		return false
	}

	m['{'] = '}'
	m['['] = ']'
	m['('] = ')'

	fmt.Println("map: ", m)

	for _, c := range s {
		switch c {
		case '{', '[', '(':
			st.push(c)
		case '}', ']', ')':
			if st.IsEmpty() {
				return false
			}
			if m[st.pop()] != c {
				return false
			}
		}
	}
	fmt.Println("Stack: ", st)
	return st.IsEmpty()
}

func main() {
	// Call isBalanced() with test cases here
	s := "{[()]}"
	b := isBalanced(s)
	fmt.Println(b)
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) push(b rune) {
	*s = append(*s, b)
}

func (s *Stack) pop() rune {
	var b rune
	if s.IsEmpty() {
		return b
	} else {
		n := len(*s)
		b = (*s)[n-1]
		*s = (*s)[:n-1]
		return b
	}
}
