/*
 * @lc app=leetcode id=389 lang=golang
 *
 * [389] Find the Difference
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	ch := t[len(t)-1]
	for i := 0; i < len(s); i++ {
		ch ^= s[i]
		ch ^= t[i]
	}
	return ch
}

// @lc code=end

