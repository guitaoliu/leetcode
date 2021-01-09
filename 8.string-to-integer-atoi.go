package Leetcode

import "math"

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	res := 0
	started := false
	neg := false
	for i := 0; i < len(s); i++ {
		if !started {
			switch {
			case s[i] == ' ':
				continue
			case s[i] == '+':
				started = true
			case s[i] == '-':
				started = true
				neg = true
			case s[i] >= '0' && s[i] <= '9':
				started = true
				res = res*10 + int(s[i]-'0')
			default:
				return res
			}
		} else {
			if s[i] >= '0' && s[i] <= '9' {
				res = res*10 + int(s[i]-'0')
				if neg && -res < math.MinInt32 {
					return math.MinInt32
				}
				if !neg && res > math.MaxInt32 {
					return math.MaxInt32
				}
			} else {
				break
			}
		}
	}
	if neg {
		return -res
	}
	return res
}

// @lc code=end
