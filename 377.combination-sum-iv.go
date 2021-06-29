/*
 * @lc app=leetcode id=377 lang=golang
 *
 * [377] Combination Sum IV
 */

// @lc code=start
func combinationSum4(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i-num >= 0 {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}

// @lc code=end

