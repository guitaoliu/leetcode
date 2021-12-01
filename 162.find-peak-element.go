/*
 * @lc app=leetcode id=162 lang=golang
 *
 * [162] Find Peak Element
 */

// @lc code=start
func findPeakElement(nums []int) int {
	left, right := 0, len(nums) - 1
	for left < right {
		mid := left + (right - left) >> 1
		if nums[mid] < nums[mid+1]{
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func findPeakElementBrute(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 && nums[i] > nums[i+1] {
			return i
		}
		if i == len(nums)-1 && nums[i] > nums[i-1] {
			return i
		}
		if i > 0 && i < len(nums)-1 && nums[i] > nums[i-1] && nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

// @lc code=end

