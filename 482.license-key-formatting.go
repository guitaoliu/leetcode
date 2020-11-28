/*
 * @lc app=leetcode id=482 lang=golang
 *
 * [482] License Key Formatting
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	cnt := 0
	sByte := []byte(S)
	res := make([]byte, 0)
	for i := len(sByte) - 1; i >= 0; i-- {
		if sByte[i] == '-' {
			continue
		}
		res = append(res, sByte[i])
		cnt++
		if cnt == K {
			res = append(res, '-')
			cnt = 0
		}
	}
	if cnt == 0 && len(res) > 0 {
		res = res[:len(res)-1]
	}

	return strings.ToUpper(string(reverse(res)))
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// @lc code=end

