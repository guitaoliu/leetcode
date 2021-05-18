/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 */

// @lc code=start
func maxProduct(words []string) int {
	if len(words) == 0 {
		return 0
	}
	bit := make([]int, len(words))
	for i, word := range words {
		for j := 0; j < len(word); j++ {
			bit[i] |= 1 << (word[j] - 'a')
		}
	}
	res := 0
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if bit[i]&bit[j] == 0 && len(words[i])*len(words[j]) > res {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}

// @lc code=end

