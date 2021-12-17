/*
 * @lc app=leetcode id=343 lang=golang
 *
 * [343] Integer Break
 */

// @lc code=start
func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[j], j)*(i-j))
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreakMath(n int) int {
	res := 0
	for i := 2; i <= n; i++ {
		dividen := n / i
		remain := n % i
		tmp := 1
		for j := 0; j < i; j++ {
			if remain > 0 {
				tmp *= dividen + 1
			} else {
				tmp *= dividen
			}
			remain--
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

// @lc code=end

