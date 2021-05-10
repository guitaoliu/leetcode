/*
 * @lc app=leetcode id=289 lang=golang
 *
 * [289] Game of Life
 */

// @lc code=start
// 2 for die
// 3 for new ones
func gameOfLife(board [][]int) {
	m, n := len(board), len(board[0])
	getNeighbors := func(i, j int) int {
		cnt := 0
		for ii := i - 1; ii <= i+1; ii++ {
			for jj := j - 1; jj <= j+1; jj++ {
				if ii == i && jj == j {
					continue
				}
				if ii >= 0 && ii < m &&
					jj >= 0 && jj < n &&
					(board[ii][jj] == 1 || board[ii][jj] == 2) {
					cnt++
				}
			}
		}
		return cnt
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := getNeighbors(i, j)
			if board[i][j] == 1 && (num < 2 || num > 3) {
				board[i][j] = 2
			}
			if board[i][j] == 0 && num == 3 {
				board[i][j] = 3
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			}
			if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

// @lc code=end

