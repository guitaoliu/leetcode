/*
 * @lc app=leetcode id=59 lang=golang
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	left, right, top, bottom := 0, n-1, 0, n-1
	cnt := 1
	for cnt < n*n {
		for idx := left; idx < right; idx++ {
			res[top][idx] = cnt
			cnt++
		}
		for idx := top; idx < bottom; idx++ {
			res[idx][right] = cnt
			cnt++
		}
		for idx := right; idx > left; idx-- {
			res[bottom][idx] = cnt
			cnt++
		}
		for idx := bottom; idx > top; idx-- {
			res[idx][left] = cnt
			cnt++
		}
		left, right, top, bottom = left+1, right-1, top+1, bottom-1
	}
	if top == bottom && left == right {
		res[top][left] = cnt
	}
	return res
}

// @lc code=end

