/*
 * @lc app=leetcode id=416 lang=golang
 *
 * [416] Partition Equal Subset Sum
 */

// @lc code=start
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	dp := make([]bool, sum/2+1)
	for i := 0; i <= sum/2; i++ {
		dp[i] = (nums[0] == i)
	}

	for i := 1; i < len(nums); i++ {
		for j := sum / 2; j > nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[sum/2]
}

// @lc code=end

