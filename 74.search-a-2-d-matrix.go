/*
 * @lc app=leetcode id=74 lang=golang
 *
 * [74] Search a 2D Matrix
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	left, right := 0, len(matrix)*len(matrix[0])-1
	n := len(matrix[0])
	for left <= right {
		mid := left + (right-left)>>1
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end

