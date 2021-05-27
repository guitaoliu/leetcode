/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
func countBits(num int) []int {
	res := make([]int, num+1)
	for i := 1; i < len(res); i++ {
		res[i] += res[i&(i-1)] + 1
	}
	return res
}

// @lc code=end

