/*
 * @lc app=leetcode id=396 lang=golang
 *
 * [396] Rotate Function
 */

// @lc code=start
func maxRotateFunction(nums []int) int {
	sum := 0
	tmp := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		tmp += i * nums[i]
	}
	n := len(nums)
	res := tmp
	for i := 0; i < n; i++ {
		tmp += sum - n*nums[n-i-1]
		res = max(tmp, res)
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

