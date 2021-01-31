/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	if m == 0 || n == 0 {
		return nil
	}
	length := m * n
	res := make([]int, length)
	cnt := 0
	left, right, top, bottom := 0, n-1, 0, m-1
	for left <= right && top <= bottom {
		for idx := left; idx < right && cnt < length; idx++ {
			res[cnt] = matrix[top][idx]
			cnt++
		}
		for idx := top; idx < bottom && cnt < length; idx++ {
			res[cnt] = matrix[idx][right]
			cnt++
		}
		for idx := right; idx > left && cnt < length; idx-- {
			res[cnt] = matrix[bottom][idx]
			cnt++
		}
		for idx := bottom; idx > top && cnt < length; idx-- {
			res[cnt] = matrix[idx][left]
			cnt++
		}
		if left == right && top == bottom && cnt < length {
			res[cnt] = matrix[left][top]
		}
		left, right, top, bottom = left+1, right-1, top+1, bottom-1
	}
	return res
}

// @lc code=end

