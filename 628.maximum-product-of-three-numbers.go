/*
 * @lc app=leetcode id=628 lang=golang
 *
 * [628] Maximum Product of Three Numbers
 */

// @lc code=start
func maximumProduct(nums []int) int {
	sort.Ints(nums)
	num1 := nums[len(nums)-1] * nums[len(nums)-2] * nums[len(nums)-3]
	num2 := nums[0] * nums[1] * nums[len(nums)-1]
	if num1 > num2 {
		return num1
	}
	return num2
}

// @lc code=end

