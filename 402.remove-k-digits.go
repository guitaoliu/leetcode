/*
 * @lc app=leetcode id=402 lang=golang
 *
 * [402] Remove K Digits
 */

// @lc code=start
func removeKdigits(num string, k int) string {
	if k == len(num) {
		return "0"
	}
	res := []byte{}
	for i := 0; i < len(num); i++ {
		ch := num[i]
		for k > 0 && len(res) > 0 && ch < res[len(res)-1] {
			res = res[:len(res)-1]
			k--
		}
		res = append(res, ch)
	}
	res = res[:len(res)-k]

	for len(res) > 1 && res[0] == '0' {
		res = res[1:]
	}

	return string(res)
}

// @lc code=end

