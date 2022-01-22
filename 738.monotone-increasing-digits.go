/*
 * @lc app=leetcode id=738 lang=golang
 *
 * [738] Monotone Increasing Digits
 */

// @lc code=start
func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))
	i := 1
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	if i < len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		i++
		for ; i < len(s); i++ {
			s[i] = '9'
		}
	}
	res, _ := strconv.Atoi(string(s))
	return res
}

// @lc code=end

