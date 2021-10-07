/*
 * @lc app=leetcode id=467 lang=golang
 *
 * [467] Unique Substrings in Wraparound String
 */

// @lc code=start
func findSubstringInWraproundString(p string) int {
	dp := make([]int, 26)
	maxContinuous := 0
	for i := 0; i < len(p); i++ {
		if i > 0 && isContinuous(p[i-1], p[i]) {
			maxContinuous++
		} else {
			maxContinuous = 1
		}
		dp[p[i]-'a'] = max(dp[p[i]-'a'], maxContinuous)
	}
	ans := 0
	for i := 0; i < len(dp); i++ {
		ans += dp[i]
	}
	return ans
}

func isContinuous(a, b byte) bool {
	if a == 'z' {
		return b == 'a'
	}
	return a+1 == b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

