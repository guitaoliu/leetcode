/*
 * @lc app=leetcode id=409 lang=golang
 *
 * [409] Longest Palindrome
 */

// @lc code=start
func longestPalindrome(s string) int {
	m := make(map[byte]int, 52)
	sByte := []byte(s)
	for i := 0; i < len(sByte); i++ {
		m[sByte[i]]++
	}
	odd, even := 0, 0
	flag := 0
	for _, v := range m {
		if v > 0 {
			if v%2 == 0 {
				even += v
			} else {
				odd += v - 1
				flag = 1
			}
		}
	}
	return odd + even + flag
}

// @lc code=end

