package main

import "fmt"

func main() {
	fmt.Println(checkPalindrome("abab"))
	fmt.Println(checkPalindrome("aba"))
	fmt.Println(checkPalindrome("madam"))
}

func partition(s string) [][]string {
	res := make([][]string, 0)

	return res
}

func checkPalindrome(s string) bool {
	if len(s) == 1 {
		return true
	}

	l := 0
	r := len(s) - 1

	for l < r {
		if s[l] == s[r] {
			l++
			r--
		} else {
			return false
		}
	}
	return true
}
