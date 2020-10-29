/*
 * @lc app=leetcode id=168 lang=golang
 *
 * [168] Excel Sheet Column Title
 */

// @lc code=start
func convertToTitle(n int) string {
	res := make([]byte, 0)
	for n > 0 {
		res = append(res, 'A'+byte((n-1)%26))
		n = (n - 1) / 26
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return string(res)
}

// @lc code=end

