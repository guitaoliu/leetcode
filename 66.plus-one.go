/*
 * @lc app=leetcode id=66 lang=golang
 *
 * [66] Plus One
 */

// @lc code=start
func plusOne(digits []int) []int {
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i]+c > 9 {
			digits[i] = (digits[i] + c) % 10
			c = 1
		} else {
			digits[i] = digits[i] + c
			c = 0
		}
	}
	if c == 1 {
		newDigits := make([]int, 1)
		newDigits[0] = 1
		newDigits = append(newDigits, digits...)
		return newDigits
	}
	return digits
}

// @lc code=end

