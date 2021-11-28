/*
 * @lc app=leetcode id=198 lang=golang
 *
 * [198] House Robber
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	a, b := nums[0], max(nums[0], nums[1])
	res := b
	for i := 2; i < len(nums); i++ {
		res = max(a+nums[i], b)
		a, b = b, res
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

