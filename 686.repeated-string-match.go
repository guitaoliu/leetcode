/*
 * @lc app=leetcode id=686 lang=golang
 *
 * [686] Repeated String Match
 */

// @lc code=start
func repeatedStringMatch(a string, b string) int {
	exist := make([]bool, 26)
	for _, c := range a {
		exist[c-'a'] = true
	}
	for _, c := range b {
		if !exist[c-'a'] {
			return -1
		}
	}

	cnt := len(b) / len(a)
	str := strings.Repeat(a, cnt)
	for i := 0; i < 3; i++ {
		if strings.Contains(str, b) {
			return cnt + i
		}
		str += a
	}
	return -1
}

// @lc code=end

