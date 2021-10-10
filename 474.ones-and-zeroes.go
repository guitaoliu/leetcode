/*
 * @lc app=leetcode id=474 lang=golang
 *
 * [474] Ones and Zeroes
 */

// @lc code=start
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zero := strings.Count(s, "0")
		one := len(s) - zero

		if zero > m || one > n {
			continue
		}

		for i := m; i >= zero; i-- {
			for j := n; j >= one; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zero][j-one])
			}
		}
	}
	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

