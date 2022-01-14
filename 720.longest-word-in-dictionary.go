/*
 * @lc app=leetcode id=720 lang=golang
 *
 * [720] Longest Word in Dictionary
 */

// @lc code=start
func longestWord(words []string) string {
	sort.Strings(words)
	m := make(map[string]bool)
	res := ""
	for _, word := range words {
		if len(word) == 1 || m[word[:len(word)-1]] {
			if len(word) > len(res) {
				res = word
			}
			m[word] = true
		}
	}
	return res
}

func longestWordHashMap(words []string) string {
	maxLength := 0
	m := make(map[int]map[string]struct{})
	for _, word := range words {
		if _, ok := m[len(word)]; !ok {
			m[len(word)] = make(map[string]struct{})
		}
		m[len(word)][word] = struct{}{}
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}
	for maxLength > 0 {
		list := []string{}
		for word := range m[maxLength] {
			l := maxLength - 1
			for l > 0 {
				if wordList, ok := m[l]; !ok {
					break
				} else {
					if _, ok := wordList[word[:l]]; !ok {
						break
					}
				}
				l--
			}
			if l == 0 {
				list = append(list, word)
			}
		}
		if len(list) > 0 {
			sort.Strings(list)
			return list[0]
		}
		maxLength--
	}
	return ""
}

// @lc code=end

