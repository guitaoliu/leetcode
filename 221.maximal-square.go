/*
 * @lc app=leetcode id=221 lang=golang
 *
 * [221] Maximal Square
 */

// @lc code=start
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	res := 0
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if res == 0 && dp[i][j] == 1 {
				res = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res * res
}

func min(nums ...int) int {
	minimum := nums[0]
	for _, v := range nums[1:] {
		if v < minimum {
			minimum = v
		}
	}
	return minimum
}

// @lc code=end

