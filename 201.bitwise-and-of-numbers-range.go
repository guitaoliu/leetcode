/*
 * @lc app=leetcode id=201 lang=golang
 *
 * [201] Bitwise AND of Numbers Range
 */

// @lc code=start
func rangeBitwiseAnd(left int, right int) int {
	for right > left {
		right &= right - 1
	}
	return right
}

// @lc code=end

