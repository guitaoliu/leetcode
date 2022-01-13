/*
 * @lc app=leetcode id=712 lang=golang
 *
 * [712] Minimum ASCII Delete Sum for Two Strings
 */

// @lc code=start
func minimumDeleteSum(s1 string, s2 string) int {
	dp := make([][]int, len(s1)+1)
	for i := 0; i <= len(s1); i++ {
		dp[i] = make([]int, len(s2)+1)
	}
	cnt := 0
	for _, ch := range s1 {
		cnt += int(ch)
	}
	for _, ch := range s2 {
		cnt += int(ch)
	}
	for i := 1; i <= len(s1); i++ {
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + int(s1[i-1])
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return cnt - 2*dp[len(s1)][len(s2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

