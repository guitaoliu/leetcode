/*
 * @lc app=leetcode id=371 lang=golang
 *
 * [371] Sum of Two Integers
 */

// @lc code=start
func getSum(a int, b int) int {
	sum := a ^ b
	carry := (a & b) << 1
	for carry != 0 {
		a := sum
		b := carry
		sum = a ^ b
		carry = (a & b) << 1
	}
	return sum
}

// @lc code=end

