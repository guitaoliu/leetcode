/*
 * @lc app=leetcode id=464 lang=golang
 *
 * [464] Can I Win
 */

// @lc code=start
func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger > desiredTotal {
		return true
	}
	if (1+maxChoosableInteger)*maxChoosableInteger/2 < desiredTotal {
		return false
	}

	m := make([]int, 1<<maxChoosableInteger)
	var trace func(int, int) bool
	trace = func(state, remain int) bool {
		if m[state] != 0 {
			return true
		}
		for i := 1; i <= maxChoosableInteger; i++ {
			cur := 1 << (i - 1)
			if state&cur == 0 {
				if i >= remain || !trace(cur|state, remain-i) {
					m[state] = 1
					return true
				}
			}
		}
		return m[state] == 1
	}
	return trace(0, desiredTotal)
}

// @lc code=end

