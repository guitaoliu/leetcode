/*
 * @lc app=leetcode id=674 lang=golang
 *
 * [674] Longest Continuous Increasing Subsequence
 */

// @lc code=start
func findLengthOfLCIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := 0
	tmp := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			tmp++
		} else {
			tmp = 0
		}
		if tmp > res {
			res = tmp
		}
	}
	return res + 1
}

// @lc code=end

