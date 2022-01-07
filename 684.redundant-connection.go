/*
 * @lc app=leetcode id=684 lang=golang
 *
 * [684] Redundant Connection
 */

// @lc code=start
func findRedundantConnection(edges [][]int) []int {
	parents := make([]int, len(edges)+1)
	for i := range parents {
		parents[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if parents[i] != i {
			return find(parents[i])
		}
		return i
	}

	union := func(a, b int) bool {
		x, y := find(a), find(b)
		if x == y {
			return false
		}
		parents[x] = y
		return true
	}

	for _, edge := range edges {
		if !union(edge[0], edge[1]) {
			return edge
		}
	}
	return nil
}

// @lc code=end

