package main

import "fmt"

func main() {
	similarWords := [][]string{
		{"very", "so"},
		{"love", "adore"},
		{"really", "very"},
		{"leetcode", "codesignal"},
		{"apples", "oranges"},
		{"like", "adore"},
		// {"code", "program"},
		// {"play", "playing"},
		// {"football", "soccer"},
		// {"like", "enjoy"},
		// {"coffee", "tea"},
	}

	sentence1 := []string{"I", "really", "love", "leetcode", "and", "apples"}
	sentence2 := []string{"I", "so", "like", "codesignal", "and", "oranges"}

	if areSentenceSimilar(sentence1, sentence2, similarWords) {
		fmt.Println("The sentences are similar.")
	} else {
		fmt.Println("The sentences are not similar.")
	}

}

func areSentenceSimilar(s1 []string, s2 []string, similarWords [][]string) bool {

	fmt.Println()

	if len(s1) != len(s2) {
		return false
	}

	similarWordSet := make(map[string]bool)

	for _, pair := range similarWords {
		if len(pair) == 2 { // Ensure each pair has exactly two words
			similarWordSet[pair[0]] = true
			similarWordSet[pair[1]] = true
		}
	}

	fmt.Println("Similar words set ", similarWordSet)
	fmt.Println()

	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	for i := 0; i < len(s1) && i < len(s2); i++ {
		if s1[i] != s2[i] {
			fmt.Println("s1[i]: ", s1[i])
			fmt.Println("s2[i]: ", s2[i])
			if similarWordSet[s1[i]] && similarWordSet[s2[i]] {
				continue
			} else {
				fmt.Println("i value: ", i)
				return false
			}
		}
	}

	return true
}
