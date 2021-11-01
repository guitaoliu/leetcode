/*
 * @lc app=leetcode id=540 lang=golang
 *
 * [540] Single Element in a Sorted Array
 */

// @lc code=start

func singleNonDuplicate(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		rightIsEven := (right-mid)%2 == 0
		if nums[mid] == nums[mid+1] {
			if rightIsEven {
				left = mid + 2
			} else {
				right = mid - 1
			}
		} else if nums[mid] == nums[mid-1] {
			if rightIsEven {
				right = mid - 2
			} else {
				left = mid + 1
			}
		} else {
			return nums[mid]
		}
	}
	return nums[left]
}

func singleNonDuplicateON(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}
	return res
}

// @lc code=end

