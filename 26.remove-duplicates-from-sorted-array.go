/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[last] == nums[finder] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		last++
		nums[last] = nums[finder]
	}
	return last + 1
}

// @lc code=end

