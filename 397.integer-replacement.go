/*
 * @lc app=leetcode id=397 lang=golang
 *
 * [397] Integer Replacement
 */

// @lc code=start
func integerReplacementBFS(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return 1 + integerReplacement(n>>1)
	} else {
		return 1 + min(integerReplacement(n+1), integerReplacement(n-1))
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func integerReplacement(n int) int {
	res := 0
	for n != 1 {
		if n&1 == 0 {
			n >>= 1
		} else {
			if n&2 == 0 || n == 3 {
				n -= 1
			} else {
				n += 1
			}
		}
		res += 1
	}
	return res
}

// @lc code=end

