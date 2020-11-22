/*
 * @lc app=leetcode id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	gi, si := 0, 0
	for gi < len(g) && si < len(s) {
		if s[si] >= g[gi] {
			gi++
			si++
			res++
		} else {
			si++
		}
	}
	return res
}

// @lc code=end

