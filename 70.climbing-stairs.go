/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairsDP(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func climbStairs(n int) int {
	a, b := 1, 1
	res := 1
	for i := 2; i < n+1; i++ {
		res = a + b
		a, b = b, res
	}
	return res
}

// @lc code=end

