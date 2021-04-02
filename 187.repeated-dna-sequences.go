/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	if len(s) < 11 {
		return []string{}
	}
	m := make(map[string]int, 0)
	res := []string{}
	for i := 0; i < len(s)-9; i++ {
		if m[s[i:i+10]] == 1 {
			res = append(res, s[i:i+10])
		}
		m[s[i:i+10]]++
	}
	return res
}

// @lc code=end

