/*
 * @lc app=leetcode id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if s[i] != t[j] {
			j++
		} else {
			i++
			j++
		}
	}
	return i == len(s)
}

// @lc code=end

