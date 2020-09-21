/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	tmp := 0
	for x != 0 {
		tmp = tmp*10 + x%10
		x /= 10
	}
	if tmp > 1<<31-1 || tmp < -(1<<31) {
		return 0
	}
	return tmp
}

// @lc code=end

