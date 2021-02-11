/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	rowFlag := false
	columnFlag := false
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				if i == 0 {
					rowFlag = true
				}
				if j == 0 {
					columnFlag = true
				}
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if rowFlag {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[0][j] = 0
		}
	}
	if columnFlag {
		for i := 1; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

// @lc code=end

