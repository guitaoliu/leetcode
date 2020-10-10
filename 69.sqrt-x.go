/*
 * @lc app=leetcode id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	left, right, res := 1, x, 0
	mid := 0
	for left <= right {
		mid = left + ((right - left) / 2)
		if mid*mid < x {
			left = mid + 1
			res = mid
		} else if mid*mid == x {
			return mid
		} else {
			right = mid - 1
		}
	}
	return res
}

// @lc code=end

