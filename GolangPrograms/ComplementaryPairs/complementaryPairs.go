package main

import "fmt"

func main() {
	stringData := []string{"abc", "abcd", "bc", "adc"}
	stringData2 := []string{"ball", "all", "call", "bal"}
	fmt.Println(countComplementaryPairs(stringData))
	fmt.Println()
	fmt.Println(countComplementaryPairs(stringData2))
}

func countComplementaryPairs(stringData []string) int {
	pairs := [][]string{}

	for i := 0; i < len(stringData); i++ {
		for j := i + 1; j < len(stringData); j++ {
			if checkPalindrome(stringData[i], stringData[j]) {
				pairs = append(pairs, []string{stringData[i], stringData[j]})
			}
		}
	}
	fmt.Println("Pairs found: ", pairs)
	return len(pairs)
}

func checkPalindrome(s1, s2 string) bool {
	freq := make(map[rune]int)
	oddCount := 0

	for _, ch := range s1 {
		freq[ch]++
	}

	for _, ch := range s2 {
		freq[ch]++
	}

	for _, val := range freq {
		if val%2 != 0 {
			oddCount++
		}
	}

	return oddCount <= 1
}
