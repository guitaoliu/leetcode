/*
 * @lc app=leetcode id=434 lang=golang
 *
 * [434] Number of Segments in a String
 */

// @lc code=start
func countSegments(s string) int {
	res := 0
	for idx, v := range s {
		if (idx == 0 || s[idx-1] == ' ') && v != ' ' {
			res++
		}
	}
	return res
}

// @lc code=end

