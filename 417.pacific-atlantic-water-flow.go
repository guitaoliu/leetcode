/*
 * @lc app=leetcode id=417 lang=golang
 *
 * [417] Pacific Atlantic Water Flow
 */

// @lc code=start

func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	p, a := make([][]bool, m), make([][]bool, m)
	for i := 0; i < m; i++ {
		p[i] = make([]bool, n)
		a[i] = make([]bool, n)
	}
	res := [][]int{}
	var dfs func(i, j, h int, visited *[][]bool)
	dfs = func(i, j, h int, visited *[][]bool) {
		if i < 0 || i >= m || j < 0 || j >= n {
			return
		}
		if heights[i][j] < h || (*visited)[i][j] {
			return
		}
		(*visited)[i][j] = true

		dfs(i+1, j, heights[i][j], visited)
		dfs(i-1, j, heights[i][j], visited)
		dfs(i, j+1, heights[i][j], visited)
		dfs(i, j-1, heights[i][j], visited)
	}

	for i := 0; i < m; i++ {
		dfs(i, 0, heights[i][0], &p)
		dfs(i, n-1, heights[i][n-1], &a)
	}
	for j := 0; j < n; j++ {
		dfs(0, j, heights[0][j], &p)
		dfs(m-1, j, heights[m-1][j], &a)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if p[i][j] && a[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

// @lc code=end

