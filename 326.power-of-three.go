/*
 * @lc app=leetcode id=326 lang=golang
 *
 * [326] Power of Three
 */

// @lc code=start
func isPowerOfThree(n int) bool {
	return n > 0 && 1162261467%n == 0
}

// @lc code=end

