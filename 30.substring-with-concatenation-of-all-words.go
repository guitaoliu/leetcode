/*
 * @lc app=leetcode id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	m := make([]int, 26)
	wordMap := make(map[string]int)
	for _, word := range words {
		for _, ch := range word {
			m[ch-'a']--
		}
		wordMap[word]++
	}
	wordLen := len(words[0])
	wordsLen := len(words) * wordLen
	left := 0
	res := []int{}
	for right := 0; right < len(s); right++ {
		m[s[right]-'a']++
		for m[s[right]-'a'] > 0 {
			m[s[left]-'a']--

		}
		if left <= len(s)-wordsLen {
			if right-left+1 == wordsLen {
				flag := true
				used := make(map[string]int)
				for i := 0; i < len(words); i++ {
					word := s[left+wordLen*i : left+wordLen*(i+1)]
					if val, ok := wordMap[word]; !ok {
						flag = false
						break
					} else {
						used[word]++
						if used[word] > val {
							flag = false
							break
						}
					}
				}
				if flag {
					res = append(res, left)
				}
			}
		}
	}
	return res
}

// @lc code=end

