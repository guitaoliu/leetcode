/*
 * @lc app=leetcode id=647 lang=golang
 *
 * [647] Palindromic Substrings
 */

// @lc code=start
func countSubstrings(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt += count(s, i, i)
		cnt += count(s, i, i+1)
	}
	return cnt
}

func count(s string, i, j int) int {
	cnt := 0
	for i >= 0 && j < len(s) && s[i] == s[j] {
		cnt++
		i--
		j++
	}
	return cnt
}

func countSubstringsDP(s string) int {
	cnt := 0
	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
		cnt++
	}
	for i := 0; i < len(dp)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 1
			cnt++
		}
	}
	for i := len(s) - 3; i >= 0; i-- {
		for j := i + 2; j < len(dp); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] == 1 {
					cnt++
				}
			}
		}
	}
	return cnt
}

// @lc code=end

