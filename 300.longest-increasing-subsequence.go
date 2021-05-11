/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}

func lengthOfLISDp(nums []int) int {
	dp := make([]int, len(nums)+1)
	dp[0] = 0
	res := 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += 1
		res = max(dp[i], res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lengthOfLISBacktrace(nums []int) int {
	res := 1
	var backtrace func(int, int, int)
	backtrace = func(idx, last, length int) {
		for i := idx; i < len(nums); i++ {
			if nums[i] > last {
				if length+1 > res {
					res = length + 1
				}
				backtrace(i+1, nums[i], length+1)
			}
		}
	}
	for i := 0; i < len(nums); i++ {
		backtrace(i, nums[i], 1)
	}
	return res
}

// @lc code=end

