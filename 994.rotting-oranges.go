/*
 * @lc app=leetcode id=994 lang=golang
 *
 * [994] Rotting Oranges
 */

// @lc code=start
var directions = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func orangesRotting(grid [][]int) int {
	queue := [][2]int{}
	cnt := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			switch grid[i][j] {
			case 1:
				cnt++
			case 2:
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	level := -1
	for len(queue) > 0 {
		cur := queue
		queue = nil
		for _, pos := range cur {
			for _, dir := range directions {
				x, y := pos[0]+dir[0], pos[1]+dir[1]
				if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
					continue
				}
				if grid[x][y] == 1 {
					queue = append(queue, [2]int{x, y})
					grid[x][y] = 2
					cnt--
				}
			}
		}
		level++
	}
	if cnt == 0 {
		if level == -1 {
			return 0
		}
		return level
	}
	return -1
}

// @lc code=end

