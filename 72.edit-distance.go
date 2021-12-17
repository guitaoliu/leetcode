/*
 * @lc app=leetcode id=72 lang=golang
 *
 * [72] Edit Distance
 */

// @lc code=start
func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 || len(word2) == 0 {
		return len(word1) + len(word2)
	}
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]))
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

