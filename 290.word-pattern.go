/*
 * @lc app=leetcode id=290 lang=golang
 *
 * [290] Word Pattern
 */

// @lc code=start
func wordPattern(pattern string, s string) bool {
	strList := strings.Split(s, " ")
	patterns := []byte(pattern)
	if pattern == "" || len(pattern) != len(strList) {
		return false
	}
	strToPattern := make(map[string]byte, 0)
	patternToStr := make(map[byte]string, 0)

	for idx, v := range patterns {
		if _, ok := patternToStr[v]; ok && patternToStr[v] != strList[idx] {
			return false
		} else if _, ok := strToPattern[strList[idx]]; ok && trToPattern[strList[idx]] != v {
			return false
		} else {
			patternToStr[v] = strList[idx]
			strToPattern[strList[idx]] = v
		}
	}
	return true
}

// @lc code=end

