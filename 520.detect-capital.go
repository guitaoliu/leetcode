/*
 * @lc app=leetcode id=520 lang=golang
 *
 * [520] Detect Capital
 */

// @lc code=start
func detectCapitalUse(word string) bool {
	wordByte := []byte(word)
	if len(word) == 1 {
		return true
	}
	if isCapital(wordByte[0]) && isCapital(wordByte[1]) {
		for _, v := range wordByte[2:] {
			if !isCapital(v) {
				return false
			}
		}
	} else if isCapital(wordByte[0]) && !isCapital(wordByte[1]) {
		for _, v := range wordByte[2:] {
			if isCapital(v) {
				return false
			}
		}
	} else if !isCapital(wordByte[0]) {
		for _, v := range wordByte[1:] {
			if isCapital(v) {
				return false
			}
		}
	}
	return true
}

func isCapital(s byte) bool {
	return s >= 'A' && s <= 'Z'
}

// @lc code=end

