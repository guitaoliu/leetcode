/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	for i := 1; i < len(res); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1
	for i := len(res) - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}
	return res
}

// @lc code=end

