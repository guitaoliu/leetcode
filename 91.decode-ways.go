/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	dp := make([]int, len(s)+1)
	dp[0] = 1
	for i := 1; i < len(dp); i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 &&
			(s[i-2] == '1' || s[i-2] == '2' && (int(s[i-1]-'0') <= 6)) {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(dp)-1]
}

// @lc code=end

