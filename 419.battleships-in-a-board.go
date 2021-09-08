/*
 * @lc app=leetcode id=419 lang=golang
 *
 * [419] Battleships in a Board
 */

// @lc code=start
func countBattleships(board [][]byte) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' {
				board[i][j] = '.'
				m, n := i, j
				for m+1 < len(board) && board[m+1][j] == 'X' {
					board[m+1][j] = '.'
					m++
				}
				for n+1 < len(board[i]) && board[i][n+1] == 'X' {
					board[i][n+1] = '.'
					n++
				}
				res++
			}
		}
	}
	return res
}

func countBattleships(board [][]byte) int {
	res := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'X' &&
				(i <= 0 || board[i-1][j] == '.') &&
				(j <= 0 || board[i][j-1] == '.') {
				res++
			}
		}
	}
	return res
}

// @lc code=end

