/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	commonStr := ""
	nums := len(strs)
	for k, v := range strs[0] {
		for i := 1; i < nums; i++ {
			if v != strs[i][k] {
				return commonStr
			}
		}
		commonStr += v
	}
	return commonStr
}

// @lc code=end

