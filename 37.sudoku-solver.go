/*
 * @lc app=leetcode id=37 lang=golang
 *
 * [37] Sudoku Solver
 */

// @lc code=start
func solveSudoku(board [][]byte) {
	row, column := [9][9]bool{}, [9][9]bool{}
	block := [3][3][9]bool{}
	pos := [][2]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if num := board[i][j]; num == '.' {
				pos = append(pos, [2]int{i, j})
			} else {
				digit := num - '1'
				row[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
			}
		}
	}

	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(pos) {
			return true
		}
		i, j := pos[index][0], pos[index][1]
		for digit := 0; digit < 9; digit++ {
			if !row[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				row[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = byte('1' + digit)
				if dfs(index + 1) {
					return true
				}
				row[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}

	dfs(0)
}

// @lc code=end

