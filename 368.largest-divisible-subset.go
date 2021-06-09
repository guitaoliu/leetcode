/*
 * @lc app=leetcode id=368 lang=golang
 *
 * [368] Largest Divisible Subset
 */

// @lc code=start
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	maxSize, maxVal := 0, 0
	for i := 0; i < len(nums); i++ {
		for j, val := range nums[:i] {
			if nums[i]%val == 0 && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxSize {
			maxSize = dp[i]
			maxVal = nums[i]
		}
	}
	res := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		if maxVal%nums[i] == 0 && dp[i] == maxSize {
			res = append(res, nums[i])
			maxVal = nums[i]
			maxSize--
		}
	}
	return res
}

// @lc code=end

