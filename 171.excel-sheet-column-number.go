/*
 * @lc app=leetcode id=171 lang=golang
 *
 * [171] Excel Sheet Column Number
 */

// @lc code=start
func titleToNumber(s string) int {
	res, cur := 0, 0
	for i := 0; i < len(s); i++ {
		cur = int(s[i] - 'A' + 1)
		res = res*26 + cur
	}
	return res
}

// @lc code=end

