/*
 * @lc app=leetcode id=500 lang=golang
 *
 * [500] Keyboard Row
 */

// @lc code=start
var keyBoardMap = map[byte]int{
	'q': 1, 'w': 1, 'e': 1, 'r': 1, 't': 1, 'y': 1, 'u': 1, 'i': 1, 'o': 1, 'p': 1,
	'a': 2, 's': 2, 'd': 2, 'f': 2, 'g': 2, 'h': 2, 'j': 2, 'k': 2, 'l': 2,
	'z': 3, 'x': 3, 'c': 3, 'v': 3, 'b': 3, 'n': 3, 'm': 3}

func findWords(words []string) []string {
	res := make([]string, 0)
	for _, word := range words {
		lowerWord := strings.ToLower(word)
		row := keyBoardMap[lowerWord[0]]
		flag := true
		for _, v := range []byte(lowerWord) {
			curRow := keyBoardMap[v]
			if curRow != row {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, word)
		}
	}
	return res
}

// @lc code=end

