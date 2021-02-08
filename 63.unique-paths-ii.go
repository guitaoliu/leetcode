/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	res[0][0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 0 && res[0][i-1] != 0 {
			res[0][i] = 1
		}
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 0 && res[i-1][0] != 0 {
			res[i][0] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				res[i][j] = res[i-1][j] + res[i][j-1]
			}
		}
	}
	return res[m-1][n-1]
}

// @lc code=end

