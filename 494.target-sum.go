/*
 * @lc app=leetcode id=494 lang=golang
 *
 * [494] Target Sum
 */

// @lc code=start
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if target > sum || (target+sum)%2 == 1 {
		return 0
	}
	half := (sum - target) / 2
	dp := make([]int, half+1)
	dp[0] = 1
	for _, num := range nums {
		for i := half; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[half]
}

func findTargetSumWaysDFS(nums []int, target int) int {
	cnt := 0
	var dfs func(int, int)
	dfs = func(idx, sum int) {
		if idx == len(nums) {
			if sum == target {
				cnt++
			}
			return
		}
		dfs(idx+1, sum+nums[idx])
		dfs(idx+1, sum-nums[idx])
		return
	}
	dfs(0, 0)
	return cnt
}

// @lc code=end

