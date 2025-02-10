package main

import "fmt"

func main() {
	res := isBalanced("{[()]}")
	fmt.Println(res)
}

func isBalanced(s string) string {
	stack := []rune{}

	bracketsMatch := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack = append(stack, char)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != bracketsMatch[char] {
				return "NO"
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) == 0 {
		return "YES"
	}

	return "NO"
}
