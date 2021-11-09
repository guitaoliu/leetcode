/*
 * @lc app=leetcode id=576 lang=golang
 *
 * [576] Out of Boundary Paths
 */

// @lc code=start

var dir = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func findPathsDP(m int, n int, maxMove int, startRow int, startColumn int) int {
	dp := make([][][]int, maxMove+1)
	for i := range dp {
		dp[i] = make([][]int, m)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
		}
	}
	dp[0][startRow][startColumn] = 1
	res := 0
	for step := 0; step < maxMove; step++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[step][i][j] > 0 {
					for _, d := range dir {
						x, y := i+d[0], j+d[1]
						if x >= 0 && y >= 0 && x < m && y < n {
							dp[step+1][x][y] = (dp[step][i][j] + dp[step+1][x][y]) % (1e9 + 7)
						} else {
							res = (res + dp[step][i][j]) % (1e9 + 7)
						}
					}
				}
			}
		}
	}
	return res
}

func findPathsDFS(m int, n int, maxMove int, startRow int, startColumn int) int {
	visited := make([][][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			visited[i][j] = make([]int, maxMove+1)
			for k := 0; k <= maxMove; k++ {
				visited[i][j][k] = -1
			}
		}
	}
	return dfs(startRow, startColumn, maxMove, m, n, visited)
}

func dfs(i, j, maxMove, m, n int, visited [][][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n {
		return 1
	}
	if maxMove == 0 {
		visited[i][j][maxMove] = 0
		return 0
	}
	if visited[i][j][maxMove] >= 0 {
		return visited[i][j][maxMove]
	}
	res := 0
	for _, d := range dir {
		x := i + d[0]
		y := j + d[1]
		res += (dfs(x, y, maxMove-1, m, n, visited)) % (1e9 + 7)
	}
	visited[i][j][maxMove] = res % (1e9 + 7)
	return visited[i][j][maxMove]
}

// @lc code=end

