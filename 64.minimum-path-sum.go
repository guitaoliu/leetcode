/*
 * @lc app=leetcode id=64 lang=golang
 *
 * [64] Minimum Path Sum
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	res[0][0] = grid[0][0]
	for i := 1; i < m; i++ {
		res[i][0] = grid[i][0] + res[i-1][0]
	}
	for i := 1; i < n; i++ {
		res[0][i] = grid[0][i] + res[0][i-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			res[i][j] = min(res[i-1][j], res[i][j-1]) + grid[i][j]
		}
	}
	return res[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

