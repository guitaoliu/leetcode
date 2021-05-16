/*
 * @lc app=leetcode id=310 lang=golang
 *
 * [310] Minimum Height Trees
 */

// @lc code=start
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	out := make([]int, n)
	m := make(map[int][]int)
	for i := 0; i < n; i++ {
		m[i] = []int{}
	}
	for _, edge := range edges {
		out[edge[0]]++
		out[edge[1]]++
		m[edge[0]] = append(m[edge[0]], edge[1])
		m[edge[1]] = append(m[edge[1]], edge[0])
	}
	queue := []int{}
	for i := 0; i < n; i++ {
		if out[i] == 1 {
			queue = append(queue, i)
		}
	}
	res := []int{}
	for len(queue) > 0 {
		res = queue
		queue = nil
		for _, node := range res {
			for _, nei := range m[node] {
				out[nei]--
				if out[nei] == 1 {
					queue = append(queue, nei)
				}
			}
		}
	}
	return res
}

// @lc code=end

