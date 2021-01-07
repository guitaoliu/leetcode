/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	for i := 0; i < len(s)-1; i++ {
		l1, r1 := expand(s, i, i)
		l2, r2 := expand(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

func expand(s string, l int, r int) (int, int) {
	for ; l >= 0 && r < len(s) && s[l] == s[r]; l, r = l-1, r+1 {
	}
	return l + 1, r - 1
}

func longestPalindromePD(s string) string {
	if len(s) == 0 {
		return ""
	}
	if len(s) < 2 {
		return s
	}

	dp := make([][]bool, len(s))
	l, r := 0, 0

	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j-i+1 > r-l+1 {
				l, r = i, j
			}
		}
	}
	return s[l : r+1]
}

// @lc code=end

