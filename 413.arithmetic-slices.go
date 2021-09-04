/*
 * @lc app=leetcode id=413 lang=golang
 *
 * [413] Arithmetic Slices
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == nums[i]-nums[i-1] {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 0
		}
	}
	return sum(dp)
}

func sum(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func numberOfArithmeticSlicesO1(nums []int) int {
	res, dp := 0, 0
	for i := 1; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] == nums[i]-nums[i-1] {
			dp += 1
			res += dp
		} else {
			dp = 0
		}
	}
	return res
}

// @lc code=end

