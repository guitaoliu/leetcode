/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	m := make([]int, 256)
	for i := 0; i < len(t); i++ {
		m[t[i]]--
	}
	left, right := 0, 0
	minLength := math.MaxInt32
	res := ""
	matched := 0
	for ; right < len(s); right++ {
		m[s[right]]++
		if m[s[right]] <= 0 {
			matched++
		}

		for matched == len(t) {
			if minLength > right-left+1 {
				minLength = right - left + 1
				res = s[left : right+1]
			}
			if m[s[left]] == 0 {
				matched--
			}
			m[s[left]]--
			left++
		}
	}
	return res
}

func check(m []int) bool {
	for _, c := range m {
		if c < 0 {
			return false
		}
	}
	return true
}

// @lc code=end

