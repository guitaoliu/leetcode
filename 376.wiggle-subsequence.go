/*
 * @lc app=leetcode id=376 lang=golang
 *
 * [376] Wiggle Subsequence
 */

// @lc code=start
func wiggleMaxLength(nums []int) int {
	dpUp := make([]int, len(nums))
	dpDown := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dpUp[i], dpDown[i] = 1, 1
	}
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dpDown[j]+1 > dpUp[i] {
				dpUp[i] = dpDown[j] + 1
			}
			if nums[i] < nums[j] && dpUp[j]+1 > dpDown[i] {
				dpDown[i] = dpUp[j] + 1
			}
		}
	}
	return max(dpDown[len(nums)-1], dpUp[len(nums)-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

