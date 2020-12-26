/*
 * @lc app=leetcode id=633 lang=golang
 *
 * [633] Sum of Square Numbers
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	for a := 0; a*a <= c; a++ {
		if bSquare := c - a*a; isSquareable(bSquare) {
			return true
		}
	}
	return false
}

func isSquareable(n int) bool {
	ix := int(math.Sqrt(float64(n)))
	return ix*ix == n
}

// @lc code=end

