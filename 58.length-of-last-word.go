/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */

// @lc code=start
func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	cnt, ptr := 0, len(s)-1
	for s[ptr] == byte(' ') {
		ptr--
		if ptr < 0 {
			return 0
		}
	}
	for s[ptr] != byte(' ') {
		cnt++
		ptr--
		if ptr < 0 {
			return cnt
		}
	}
	return cnt
}

// @lc code=end

