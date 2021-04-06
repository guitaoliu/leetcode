/*
 * @lc app=leetcode id=207 lang=golang
 *
 * [207] Course Schedule
 */

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	in := make([]int, numCourses)
	for _, v := range prerequisites {
		in[v[0]]++
		edges[v[1]] = append(edges[v[1]], v[0])
	}
	startNodes := make([]int, 0)
	for i, v := range in {
		if v == 0 {
			startNodes = append(startNodes, i)
		}
	}
	for i := 0; i < len(startNodes); i++ {
		v := startNodes[i]
		es := edges[v]
		for _, e := range es {
			in[e]--
			if in[e] == 0 {
				startNodes = append(startNodes, e)
			}
		}
	}
	return len(startNodes) == numCourses
}

// @lc code=end

