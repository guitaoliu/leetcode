/*
 * @lc app=leetcode id=566 lang=golang
 *
 * [566] Reshape the Matrix
 */

// @lc code=start
func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums)*len(nums[0]) != r*c {
		return nums
	}
	m, n := 0, 0
	update := func(m, n int) (int, int) {
		n++
		if n == len(nums[m]) {
			m, n = m+1, 0
		}
		return m, n
	}
	res := make([][]int, r)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, c)
		for j := 0; j < len(res[i]); j++ {
			res[i][j] = nums[m][n]
			m, n = update(m, n)
		}
	}
	return res
}

// @lc code=end

