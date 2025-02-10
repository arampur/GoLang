package main

import "fmt"

func main() {
	// text := "splunk"
	// pattern := "sp+u+k"

	// res := isMatch(text, pattern, 0, 0)
	// fmt.Println("result: ", res)
	// res = isMatch("splunk", "*pl*k", 0, 0)
	// fmt.Println("result: ", res)

	fmt.Println(isMatch("Splunk", "Spl*", 0, 0))   // True
	fmt.Println(isMatch("Splunk", "Sp+u+k", 0, 0)) // True
	fmt.Println(isMatch("Splunk", "S*l+n+", 0, 0)) // True
	fmt.Println(isMatch("Splunk", "S+l*k", 0, 0))  // True
	fmt.Println(isMatch("Splunk", "Sp+n", 0, 0))   // False

}

func isMatch(text string, pattern string, textIndex int, patternIndx int) bool {
	if patternIndx == len(pattern) {
		return textIndex == len(text)
	}

	if pattern[patternIndx] == '+' {
		if textIndex < len(text) {
			return isMatch(text, pattern, textIndex+1, patternIndx+1)
		}
	}

	if pattern[patternIndx] == '*' {
		return isMatch(text, pattern, textIndex, patternIndx+1) || (textIndex < len(text) && isMatch(text, pattern, textIndex+1, patternIndx))
	}

	if textIndex < len(text) && patternIndx < len(pattern) && pattern[patternIndx] == text[textIndex] {
		return isMatch(text, pattern, textIndex+1, patternIndx+1)
	}

	return false
}
