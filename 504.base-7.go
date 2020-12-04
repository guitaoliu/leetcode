/*
 * @lc app=leetcode id=504 lang=golang
 *
 * [504] Base 7
 */

// @lc code=start
func convertToBase7(num int) string {
	res := []byte{}
	if num < 0 {
		return "-" + convertToBase7(-num)
	}
	if num == 0 {
		return "0"
	}
	for num > 0 {
		res = append(res, byte('0'+num%7))
		num /= 7
	}
	reverse := func(s []byte) []byte {
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}
	return string(reverse(res))
}

// @lc code=end

