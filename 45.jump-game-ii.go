/*
 * @lc app=leetcode id=45 lang=golang
 *
 * [45] Jump Game II
 */

// @lc code=start
func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	canReach := 0
	step := 0
	nextStep := 0
	for i, v := range nums {
		if i+v > canReach {
			canReach = i + v
			if canReach >= len(nums)-1 {
				return step + 1
			}
		}
		if i == nextStep {
			nextStep = canReach
			step++
		}
	}
	return 0
}

func jumpDP(nums []int) int {
	dp := make([]int, len(nums))
	for i := 1; i < len(dp); i++ {
		dp[i] = math.MaxInt32
	}
	for i := 0; i < len(dp); i++ {
		for j := 0; j < nums[i]; j++ {
			next := i + j + 1
			if next >= len(nums) {
				break
			}
			dp[next] = min(dp[next], dp[i]+1)
		}
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

