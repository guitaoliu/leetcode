/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start
func characterReplacement(s string, k int) int {
	res, left, cnt := 0, 0, 0
	freq := make([]int, 26)
	for right := 0; right < len(s); right++ {
		freq[s[right]-'A']++
		cnt := max(cnt, freq[s[right]-'A'])
		for right-left+1-cnt > k {
			freq[s[left]-'A']--
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

