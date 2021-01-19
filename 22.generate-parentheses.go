/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */

// @lc code=start
func generateParenthesis(n int) []string {
	res := []string{}

	var generate func(int, int, string)
	generate = func(lIdx, rIdx int, str string) {
		if lIdx == 0 && rIdx == 0 {
			res = append(res, str)
			return
		}
		if lIdx > 0 {
			generate(lIdx-1, rIdx, str+"(")
		}
		if rIdx > 0 && lIdx < rIdx {
			generate(lIdx, rIdx-1, str+")")
		}
	}

	generate(n, n, "")
	return res
}

// @lc code=end

