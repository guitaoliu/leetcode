/*
 * @lc app=leetcode id=486 lang=golang
 *
 * [486] Predict the Winner
 */

// @lc code=start
func PredictTheWinner(nums []int) bool {
	dp := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = make([]int, len(nums))
		dp[i][i] = nums[i]
	}
	for i := len(nums) - 2; i >= 0; i-- {
		for j := i + 1; j < len(nums); j++ {
			pickI := nums[i] - dp[i+1][j]
			pickJ := nums[j] - dp[i][j-1]
			dp[i][j] = max(pickI, pickJ)
		}
	}
	return dp[0][len(nums)-1] >= 0
}

func PredictTheWinnerTrace(nums []int) bool {
	m := make([][]int, len(nums))
	for i := 0; i < len(m); i++ {
		m[i] = make([]int, len(nums))
		for j := 0; j < len(nums); j++ {
			m[i][j] = math.MinInt64
		}
	}
	var trace func(int, int) int
	trace = func(i, j int) int {
		if i == j {
			return nums[i]
		}
		if m[i][j] != math.MinInt64 {
			return m[i][j]
		}
		scoreI := nums[i] - trace(i+1, j)
		scoreJ := nums[j] - trace(i, j-1)
		return max(scoreI, scoreJ)
	}
	return trace(0, len(nums)-1) >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

