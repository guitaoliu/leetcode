/*
 * @lc app=leetcode id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}

	res := []int{}
	m := make([]int, 26)
	for i := 0; i < len(p); i++ {
		m[s[i]-'a']--
		m[p[i]-'a']++
	}
	if checkIdentical(m) {
		res = append(res, 0)
	}
	for i := 1; i < len(s)-len(p)+1; i++ {
		m[s[i-1]-'a']++
		m[s[i+len(p)-1]-'a']--
		if checkIdentical(m) {
			res = append(res, i)
		}
	}
	return res
}

func checkIdentical(m []int) bool {
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

