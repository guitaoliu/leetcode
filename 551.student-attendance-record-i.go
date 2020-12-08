/*
 * @lc app=leetcode id=551 lang=golang
 *
 * [551] Student Attendance Record I
 */

// @lc code=start
func checkRecord(s string) bool {
	cntA := 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			cntA++
			if cntA > 1 {
				return false
			}
		} else if (i < len(s)-2 && s[i] == 'L' && s[i+1] == 'L' && s[i+2] == 'L') ||
			(i > 0 && i < len(s)-1 && s[i-1] == 'L' && s[i] == 'L' && s[i+1] == 'L') ||
			(i > 1 && s[i-2] == 'L' && s[i-1] == 'L' && s[i] == 'L') {
			return false
		}
	}
	return true
}

// @lc code=end

