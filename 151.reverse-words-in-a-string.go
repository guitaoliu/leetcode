/*
 * @lc app=leetcode id=151 lang=golang
 *
 * [151] Reverse Words in a String
 */

// @lc code=start
func reverseWords(s string) string {
	sSplit := strings.Fields(s)
	for i, j := 0, len(sSplit)-1; i < j; i, j = i+1, j-1 {
		sSplit[i], sSplit[j] = sSplit[j], sSplit[i]
	}
	return strings.Join(sSplit, " ")
}

// @lc code=end

