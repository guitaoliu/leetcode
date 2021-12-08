/*
 * @lc app=leetcode id=1091 lang=golang
 *
 * [1091] Shortest Path in Binary Matrix
 */

// @lc code=start
var directions = [][]int{{0, -1}, {-1, 0}, {1, 0}, {0, 1}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 {
		return -1
	}
	if n == 1 && grid[0][0] == 0 {
		return 1
	}
	queue := [][]int{{0, 0}}
	level := 2
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		for _, pos := range curLevel {
			i, j := pos[0], pos[1]
			for _, dir := range directions {
				x, y := i+dir[0], j+dir[1]
				if x < 0 || x >= n || y < 0 || y >= n || grid[x][y] != 0 {
					continue
				}
				grid[x][y] = 2
				queue = append(queue, []int{x, y})
				if x == n-1 && y == n-1 {
					return level
				}
			}
		}
		level++
	}
	return -1
}

// @lc code=end

