package main

import "fmt"

func main() {
	matchCheck := isMatch("aa", "a")
	fmt.Println("Result:", matchCheck)
}

func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)

	fmt.Println("m:", m)
	fmt.Println("n:", n)

	dp := make([][]bool, m+1)
	fmt.Println("DP: ", dp)
	fmt.Println()
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
		fmt.Println("DP[i]: ", dp[i])
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
			fmt.Println("dp[0][i] = dp[0][i-2]", dp[0][i])
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '.' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
				fmt.Println("dp[i][j] = dp[i-1][j-1]", dp[i][j])
			} else if p[j-1] == '*' { // pa[j-1]=="*"
				dp[i][j] = dp[i][j-2]
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	fmt.Println("dp[m][n]", dp[m][n])
	return dp[m][n]
}
