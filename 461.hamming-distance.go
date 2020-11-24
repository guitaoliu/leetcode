/*
 * @lc app=leetcode id=461 lang=golang
 *
 * [461] Hamming Distance
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	res := 0
	for xor := x ^ y; xor != 0; xor &= xor - 1 {
		res++
	}
	return res
}

// @lc code=end

