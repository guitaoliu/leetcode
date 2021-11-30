/*
 * @lc app=leetcode id=977 lang=golang
 *
 * [977] Squares of a Sorted Array
 */

// @lc code=start
func sortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	res := make([]int, len(nums))
	tail := len(nums) - 1
	for i <= j {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			res[tail] = nums[i] * nums[i]
			i++
		} else {
			res[tail] = nums[j] * nums[j]
			j--
		}
		tail--
	}
	return res
}

// @lc code=end

