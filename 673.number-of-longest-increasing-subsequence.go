/*
 * @lc app=leetcode id=673 lang=golang
 *
 * [673] Number of Longest Increasing Subsequence
 */

// @lc code=start

func findNumberOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dpN := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		dpN[i] = 1
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					dpN[i] = dpN[j]
				} else if dp[j]+1 == dp[i] {
					dpN[i] += dpN[j]
				}
			}
		}
	}
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		if dp[i] > maxLen {
			maxLen = dp[i]
		}
	}
	res := 0
	for i := 0; i < len(nums); i++ {
		if dp[i] == maxLen {
			res += dpN[i]
		}
	}
	return res
}

// @lc code=end
