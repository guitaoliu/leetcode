/*
 * @lc app=leetcode id=485 lang=golang
 *
 * [485] Max Consecutive Ones
 */

// @lc code=start
func findMaxConsecutiveOnes(nums []int) int {
	res := 0
	cur := 0
	for _, v := range nums {
		if v == 1 {
			cur++
		} else {
			cur = 0
		}
		if cur > res {
			res = cur
		}
	}
	return res
}

// @lc code=end

