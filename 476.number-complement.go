/*
 * @lc app=leetcode id=476 lang=golang
 *
 * [476] Number Complement
 */

// @lc code=start
func findComplement(num int) int {
	mask := 1
	for mask <= num {
		mask <<= 1
	}
	return (mask - 1) ^ num
}

// @lc code=end

