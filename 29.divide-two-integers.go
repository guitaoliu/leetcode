/*
 * @lc app=leetcode id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	neg := -1
	if dividend >= 0 && divisor >= 0 || dividend < 0 && divisor < 0 {
		neg = 1
	}
	dividend, divisor = abs(dividend), abs(divisor)
	start, end := 0, dividend
	for start <= end {
		mid := start + (end-start)/2
		if mid*divisor == dividend {
			return mid * neg
		} else if mid*divisor > dividend {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return neg * end
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

// @lc code=end

