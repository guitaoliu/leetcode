/*
 * @lc app=leetcode id=542 lang=golang
 *
 * [542] 01 Matrix
 */

// @lc code=start

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] != 0 {
				left, top := math.MaxInt32, math.MaxInt32
				if i > 0 {
					left = mat[i-1][j] + 1
				}
				if j > 0 {
					top = mat[i][j-1] + 1
				}
				mat[i][j] = min(left, top)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if mat[i][j] != 0 {
				right, bottom := math.MaxInt32, math.MaxInt32
				if i < m-1 {
					right = mat[i+1][j] + 1
				}
				if j < n-1 {
					bottom = mat[i][j+1] + 1
				}
				mat[i][j] = min(min(mat[i][j], right), bottom)
			}
		}
	}
	return mat
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func updateMatrixDFS(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 &&
				!(i > 0 && mat[i-1][j] == 0 ||
					i < m-1 && mat[i+1][j] == 0 ||
					j > 0 && mat[i][j-1] == 0 ||
					j < n-1 && mat[i][j+1] == 0) {
				mat[i][j] = math.MaxInt32
			}
		}
	}
	var dfs func(int, int, int)
	dfs = func(i, j, d int) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if d >= mat[i][j] {
			return
		}
		if d > 0 {
			mat[i][j] = d
		}
		for _, dir := range directions {
			dfs(i+dir[0], j+dir[1], mat[i][j]+1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				dfs(i, j, -1)
			}
		}
	}
	return mat
}

func updateMatrixBFS(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	queue := [][]int{}
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				res[i][j] = -1
				queue = append(queue, []int{i, j})
			}
		}
	}
	level := 1
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		for _, node := range curLevel {
			for _, dir := range directions {
				x, y := node[0]+dir[0], node[1]+dir[1]
				if x < 0 || x >= m || y < 0 || y >= n {
					continue
				}
				if res[x][y] < 0 || res[x][y] > 0 {
					continue
				}
				res[x][y] = level
				queue = append(queue, []int{x, y})
			}
		}
		level++
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if res[i][j] == -1 {
				res[i][j] = 0
			}
		}
	}
	return res
}

// @lc code=end

