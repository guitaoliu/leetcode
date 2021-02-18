/*
 * @lc app=leetcode id=81 lang=golang
 *
 * [81] Search in Rotated Sorted Array II
 */

// @lc code=start
func search(nums []int, target int) bool {
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)>>1
		if nums[mid] == target {
			return true
		} else if nums[mid] > nums[start] {
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else if nums[mid] < nums[end] {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		} else {
			if nums[mid] == nums[start] {
				start++
			}
			if nums[mid] == nums[end] {
				end--
			}
		}
	}
	return false
}

// @lc code=end

