/*
 * @lc app=leetcode id=547 lang=golang
 *
 * [547] Number of Provinces
 */

// @lc code=start
func findCircleNum(isConnected [][]int) int {
	m := make([]int, len(isConnected))
	var dfs func(int, int)
	dfs = func(node, groupNum int) {
		if m[node] != 0 && groupNum >= m[node] {
			return
		}
		m[node] = groupNum
		for j := 0; j < len(isConnected[node]); j++ {
			if node != j && isConnected[node][j] == 1 {
				dfs(j, groupNum)
			}
		}
	}
	for i := 0; i < len(isConnected); i++ {
		dfs(i, i+1)
	}
	cnt := make(map[int]struct{})
	for i := 0; i < len(m); i++ {
		if _, ok := cnt[m[i]]; !ok {
			cnt[m[i]] = struct{}{}
		}
	}
	return len(cnt)
}

// @lc code=end

