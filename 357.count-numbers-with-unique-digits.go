/*
 * @lc app=leetcode id=357 lang=golang
 *
 * [357] Count Numbers with Unique Digits
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	res := 10
	cur := 9
	for i := 0; i < n-1; i++ {
		cur *= (9 - i)
		res += cur
	}
	return res
}

// @lc code=end

