/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	var backtracing func(int, int, string)
	backtracing = func(idx int, cnt int, result string) {
		if idx == len(s) {
			if cnt == 4 {
				res = append(res, result[1:])
			}
		} else {
			if cnt == 4 {
				return
			}
			for i := idx; i < idx+3 && i < len(s); i++ {
				if i > idx && s[idx] == '0' {
					return
				}
				v, _ := strconv.Atoi(s[idx : i+1])
				if v < 256 {
					lenResult := len(result)
					result = result + "." + s[idx:i+1]
					backtracing(i+1, cnt+1, result)
					result = result[:lenResult]
				}
			}
		}
	}
	var ss string
	backtracing(0, 0, ss)
	return res
}

// @lc code=end

