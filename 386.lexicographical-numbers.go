/*
 * @lc app=leetcode id=386 lang=golang
 *
 * [386] Lexicographical Numbers
 */

// @lc code=start
func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	var dfs func(x int)
	dfs = func(x int) {
		max := (x + 10) / 10 * 10
		for x <= n && x < max {
			res = append(res, x)
			if x*10 <= n {
				dfs(x * 10)
			}
			x++
		}
	}

	dfs(1)

	return res
}

// @lc code=end

