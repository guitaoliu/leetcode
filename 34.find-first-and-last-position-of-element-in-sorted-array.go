/*
 * @lc app=leetcode id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	low, high := 0, len(nums)-1
	res := []int{-1, -1}
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == 0 || nums[mid-1] != target {
				res[0] = mid
			}
			high = mid - 1
		}
	}
	low, high = 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		if nums[mid] > target {
			high = mid - 1
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			if mid == len(nums)-1 || nums[mid+1] != target {
				res[1] = mid
			}
			low = mid + 1
		}
	}
	return res
}

// @lc code=end

