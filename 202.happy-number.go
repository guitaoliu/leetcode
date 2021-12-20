/*
 * @lc app=leetcode id=202 lang=golang
 *
 * [202] Happy Number
 */

// @lc code=start
func isHappy(n int) bool {
	m := 0
	m |= 1 << n
	for n != 1 {
		n = sum(n)
		if 1<<n&m != 0 {
			return false
		}
		m |= 1 << n
	}
	return true
}

func sum(n int) int {
	res := 0
	for n > 0 {
		mod := n % 10
		n /= 10
		res += mod * mod
	}
	return res
}

// @lc code=end

