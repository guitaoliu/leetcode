/*
 * @lc app=leetcode id=524 lang=golang
 *
 * [524] Longest Word in Dictionary through Deleting
 */

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	pointers := make([]int, len(dictionary))
	for i := 0; i < len(s); i++ {
		for idx, word := range dictionary {
			if pointers[idx] < len(word) && s[i] == word[pointers[idx]] {
				pointers[idx]++
			}
		}
	}
	res := ""
	for i := 0; i < len(dictionary); i++ {
		if pointers[i] == len(dictionary[i]) {
			if len(res) < len(dictionary[i]) || len(res) == len(dictionary[i]) && res > dictionary[i] {
				res = dictionary[i]
			}
		}
	}
	return res
}

// @lc code=end

