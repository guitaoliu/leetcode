/*
 * @lc app=leetcode id=712 lang=golang
 *
 * [712] Minimum ASCII Delete Sum for Two Strings
 */

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	for i := len(s1) - 1; i >= 0; i-- {
		dp[i][len(s2)] = dp[i+1][len(s2)] + int(s1[i])
	}
	for j := len(s2) - 1; j >= 0; j-- {
		dp[len(s1)][j] = dp[len(s1)][j] + int(s2[j])
	}
	for i := len(s1) - 1; i >= 0; i-- {
		for j := len(s2) - 1; j >= 0; j-- {
			if s1[j] == s2[j] {
				dp[i][j] = dp[i+1][j+1]
			} else {
				dp[i][j] = min(dp[i+1][j]+int(s1[i]), dp[i][j+1]+int(s2[j]))
			}
		}
	}
	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

