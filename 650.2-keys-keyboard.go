/*
 * @lc app=leetcode id=650 lang=golang
 *
 * [650] 2 Keys Keyboard
 */

// @lc code=start
func minSteps(n int) int {
	dp := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		dp[i] = math.MaxInt64
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				dp[i] = min(dp[i], dp[j]+i/j)
				dp[i] = min(dp[i], dp[i/j]+j)
			}
		}
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

