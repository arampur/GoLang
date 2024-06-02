package main

import (
	"fmt"
)

func main() {
	arr := []string{"p", "e", "r", "f", "e", "c", "t", " ",
		"m", "a", "k", "e", "s", " ",
		"p", "r", "a", "c", "t", "i", "c", "e"}

	//output for this program should be like:
	//[p r a c t i c e   m a k e s   p e r f e c t]

	res := ReverseWords(arr)
	fmt.Println(res)
}

func ReverseWords(arr []string) []string {
	// your code goes here
	n := len(arr)

	reverseEach(arr, 0, n-1)

	wordStart := 0

	for i := 0; i < n; i++ {
		if arr[i] == " " {
			reverseEach(arr, wordStart, i-1) // to reverse word before space character
			wordStart = i + 1
		}
	}

	reverseEach(arr, wordStart, n-1)
	return arr
}

func reverseEach(arr []string, start, end int) {
	temp := ""
	for start < end {
		temp = arr[start]
		arr[start] = arr[end]
		arr[end] = temp
		start++
		end--
	}
}
