/*
 * @lc app=leetcode id=306 lang=golang
 *
 * [306] Additive Number
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	if len(num) < 3 {
		return false
	}
	var backtrace func(int, int, int, int) bool
	backtrace = func(idx, cnt, prepre, pre int) bool {
		if idx == len(num) && cnt >= 3 {
			return true
		}
		res := false
		sum := prepre + pre
		for i := idx; i < len(num); i++ {
			if i > idx && num[idx] == '0' {
				break
			}
			num, _ := strconv.Atoi(num[idx : i+1])
			if cnt >= 2 {
				if num > sum {
					break
				} else if num < sum {
					continue
				}
			}
			res = res || backtrace(i+1, cnt+1, pre, num)
		}
		return res
	}
	return backtrace(0, 0, 0, 0)
}

// @lc code=end

