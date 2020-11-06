/*
 * @lc app=leetcode id=263 lang=golang
 *
 * [263] Ugly Number
 */

// @lc code=start
func isUgly(num int) bool {
	if num == 0 {
		return false
	}
	for num%2 == 0 {
		num /= 2
	}
	for num%3 == 0 {
		num /= 3
	}
	for num%5 == 0 {
		num /= 5
	}
	return num == 1
}

// @lc code=end

