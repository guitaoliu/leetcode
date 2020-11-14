/*
 * @lc app=leetcode id=383 lang=golang
 *
 * [383] Ransom Note
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		if m[v] > 0 {
			m[v]--
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

