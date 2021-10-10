/*
 * @lc app=leetcode id=473 lang=golang
 *
 * [473] Matchsticks to Square
 */

// @lc code=start
func makesquare(matchsticks []int) bool {
	sum := 0
	for _, v := range matchsticks {
		sum += v
	}
	if sum%4 != 0 {
		return false
	}
	edgeLen := sum / 4
	sort.Slice(matchsticks, func(i, j int) bool { return matchsticks[i] > matchsticks[j] })
	edges := make([]int, 4)
	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return edges[0] == edgeLen && edges[1] == edgeLen && edges[2] == edgeLen && edges[3] == edgeLen
		}
		for i := 0; i < 4; i++ {
			if edges[i]+matchsticks[idx] <= edgeLen {
				edges[i] += matchsticks[idx]
				if dfs(idx + 1) {
					return true
				}
				edges[i] -= matchsticks[idx]
			}
		}
		return false
	}
	return dfs(0)
}

// @lc code=end

