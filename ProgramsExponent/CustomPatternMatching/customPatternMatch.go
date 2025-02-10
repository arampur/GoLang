// Custom Pattern Matching to return true in below cases
// fmt.Println(matchPattern("Splunk", "Spl*")) // True
// fmt.Println(matchPattern("Splunk", "Sp+u+k")) // True
// fmt.Println(matchPattern("Splunk", "S*l+n+")) // True
// fmt.Println(matchPattern("Splunk", "S+l*k")) // True
// fmt.Println(matchPattern("Splunk", "Sp+n")) // False
package main

import "fmt"

func main() {
	fmt.Println(isMatch("splunk", "spl*", 0, 0))   // True
	fmt.Println(isMatch("splunk", "sp+u+k", 0, 0)) // True
	fmt.Println(isMatch("splunk", "s*l+n+", 0, 0)) // True
	fmt.Println(isMatch("splunk", "s+l*k", 0, 0))  // True
	fmt.Println(isMatch("splunk", "sp+n", 0, 0))   // False
}

func isMatch(text string, pattern string, textIndex int, patternIndex int) bool {

	if patternIndex == len(pattern) {
		return textIndex == len(text)
	}

	if pattern[patternIndex] == '+' {
		if textIndex < len(text) {
			return isMatch(text, pattern, textIndex+1, patternIndex+1)
		}
		return false
	}

	if pattern[patternIndex] == '*' {
		return isMatch(text, pattern, textIndex, patternIndex+1) || (textIndex < len(text) && isMatch(text, pattern, textIndex+1, patternIndex))
	}

	if textIndex < len(text) && pattern[patternIndex] == text[textIndex] {
		return isMatch(text, pattern, textIndex+1, patternIndex+1)
	}

	return false
}
