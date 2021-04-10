/*
 * @lc app=leetcode id=210 lang=golang
 *
 * [210] Course Schedule II
 */

// @lc code=start
func findOrder(numCourses int, prerequisites [][]int) []int {
	in := make([]int, numCourses)
	edges := make([][]int, numCourses)
	for _, edge := range prerequisites {
		in[edge[0]]++
		edges[edge[1]] = append(edges[edge[1]], edge[0])
	}
	startNodes := make([]int, 0)
	for i, v := range in {
		if v == 0 {
			startNodes = append(startNodes, i)
		}
	}
	for i := 0; i < len(startNodes); i++ {
		v := startNodes[i]
		for _, e := range edges[v] {
			in[e]--
			if in[e] == 0 {
				startNodes = append(startNodes, e)
			}
		}
	}
	if len(startNodes) == numCourses {
		return startNodes
	}
	return nil
}

// @lc code=end

