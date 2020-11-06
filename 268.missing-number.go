/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res + i + 1 - nums[i]
	}
	return res
}

// @lc code=end

