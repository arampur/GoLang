package main

/*
Given two strings of uppercase letters source and target, list (in string form) a sequence of edits to convert from source to target that uses the fewest edits possible.

For example, with strings source = "ABCDEFG", and target = "ABDFFGH" we might return: ["A", "B", "-C", "D", "-E", "F", "+F", "G", "+H"
More formally, for each character C in source, we will either write the token C, which does not count as an edit; or write the token -C, which counts as an edit.

Additionally, between any token that we write, we may write +D where D is any letter, which counts as an edit.
At the end, when reading the tokens from left to right, and not including tokens prefixed with a minus-sign, the letters should spell out target (when ignoring plus-signs.)

In the example, the answer of A B -C D -E F +F G +H has total number of edits 4 (the minimum possible), and ignoring subtraction-tokens, spells out A, B, D, F, +F, G, +H which represents the string target.
If there are multiple answers, use the answer that favors removing from the source first.
*/

import (
	"fmt"
	"strings"
)

func diffBetweenTwoStrings(source, target string) []string {
	m, n := len(source), len(target)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Initialize the dp table
	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if i == m {
				dp[i][j] = n - j
			} else if j == n {
				dp[i][j] = m - i
			} else if source[i] == target[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = 1 + min(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	var result []string
	i, j := 0, 0
	for i < m && j < n {
		if source[i] == target[j] {
			result = append(result, string(source[i]))
			i++
			j++
		} else if dp[i+1][j] <= dp[i][j+1] {
			result = append(result, "-"+string(source[i]))
			i++
		} else {
			result = append(result, "+"+string(target[j]))
			j++
		}
	}

	for i < m {
		result = append(result, "-"+string(source[i]))
		i++
	}

	for j < n {
		result = append(result, "+"+string(target[j]))
		j++
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	source := "ABCDEFG"
	target := "ABDFFGH"
	result := diffBetweenTwoStrings(source, target)
	//fmt.Println(strings.Join(result, " "))

	finalString := []string{}
	for i := 0; i < len(result); i++ {
		if !strings.Contains(result[i], "-") {
			finalString = append(finalString, string(result[i]))
		}
	}
	fmt.Println(finalString)
}
