/*
 * @lc app=leetcode id=554 lang=golang
 *
 * [554] Brick Wall
 */

// @lc code=start
func leastBricks(wall [][]int) int {
	m := make(map[int]int)
	for _, row := range wall {
		pos := 0
		for i := 0; i < len(row)-1; i++ {
			pos += row[i]
			m[pos]++
		}
	}
	max := 0
	for _, v := range m {
		if v > max {
			max = v
		}
	}
	return len(wall) - max
}

// @lc code=end

