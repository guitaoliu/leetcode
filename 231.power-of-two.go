/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */

// @lc code=start
// func isPowerOfTwo(n int) bool {
// 	for n >= 2 {
// 		if n%2 == 0 {
// 			n /= 2
// 		} else {
// 			return false
// 		}
// 	}
// 	return n == 1
// }

// func isPowerOfTwo(n int) bool {
// 	return n > 0 && (1073741824%n == 0)
// }

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

// @lc code=end

