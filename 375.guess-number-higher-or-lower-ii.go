/*
 * @lc app=leetcode id=375 lang=golang
 *
 * [375] Guess Number Higher or Lower II
 */

// @lc code=start
import (
	"math"
)

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for ll := 2; ll <= n; ll++ {
		for i := 1; i <= n-ll+1; i++ {
			cur := math.MaxInt32
			// for piv := i; piv < i+ll-1; piv++ {
			for piv := i + (ll-1)/2; piv < i+ll-1; piv++ {
				cur = min(cur, piv+max(dp[i][piv-1], dp[piv+1][i+ll-1]))
			}
			dp[i][i+ll-1] = cur
		}
	}
	return dp[1][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

