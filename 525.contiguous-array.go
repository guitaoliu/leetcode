/*
 * @lc app=leetcode id=525 lang=golang
 *
 * [525] Contiguous Array
 */

// @lc code=start
func findMaxLength(nums []int) int {
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			sum--
		} else {
			sum++
		}
		if index, ok := m[sum]; ok {
			res = max(res, i-index)
		} else {
			m[sum] = i
		}
	}
	return res
}

func findMaxLengthBrute(nums []int) int {
	dp := make([]int, len(nums)+1)
	res := 0
	dp[0] = 0
	for i := 0; i < len(nums); i++ {
		dp[i+1] = dp[i] + nums[i]
		for j := i - 1; j >= 0; j-- {
			if (dp[i+1]-dp[j])*2 == i+1-j {
				res = max(res, i+1-j)
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

