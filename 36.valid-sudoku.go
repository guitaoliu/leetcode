/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	num := make([]int, len(board))
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if v := board[i][j]; v != '.' {
				if num[v-'1'] == 1 {
					return false
				}
				num[v-'1'] = 1
			}
		}
		num = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	}
	for j := 0; j < len(board); j++ {
		for i := 0; i < len(board); i++ {
			if v := board[i][j]; v != '.' {
				if num[v-'1'] == 1 {
					return false
				}
				num[v-'1'] = 1
			}
		}
		num = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
	}
	for k := 0; k < 3; k++ {
		for l := 0; l < 3; l++ {
			for i := 3 * k; i < 3*(k+1); i++ {
				for j := 3 * l; j < 3*(l+1); j++ {
					if v := board[i][j]; v != '.' {
						if num[v-'1'] == 1 {
							return false
						}
						num[v-'1'] = 1
					}
				}
			}
			num = []int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		}
	}
	return true
}

// @lc code=end

