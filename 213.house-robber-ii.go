/*
 * @lc app=leetcode id=213 lang=golang
 *
 * [213] House Robber II
 */

// @lc code=start
func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	return max(robRange(nums, 0, len(nums)-1), robRange(nums, 1, len(nums)))
}

func robRange(nums []int, start, end int) int {
	a, b := nums[start], max(nums[start], nums[start+1])
	for i := start + 2; i < end; i++ {
		tmp := b
		b = max(a+nums[i], b)
		a = tmp
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

