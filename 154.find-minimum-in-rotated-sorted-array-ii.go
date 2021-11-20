/*
 * @lc app=leetcode id=154 lang=golang
 *
 * [154] Find Minimum in Rotated Sorted Array II
 */

// @lc code=start
func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left] < nums[right] {
			return nums[left]
		}
		mid := left + (right-left)>>1
		if nums[mid] > nums[left] {
			left = mid + 1
		} else if nums[mid] == nums[left] {
			left++
		} else {
			right = mid
		}

	}
	return nums[left]
}

// @lc code=end

