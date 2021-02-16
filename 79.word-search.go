/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
var directions = [][]int{
	[]int{-1, 0},
	[]int{0, 1},
	[]int{1, 0},
	[]int{0, -1},
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	var trace func(i, j, k int) bool
	trace = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visited[i][j] = true
		defer func() {
			visited[i][j] = false
		}()
		for _, dir := range directions {
			if newI, newJ := i+dir[0], j+dir[1]; newI >= 0 && newI < len(board) && newJ >= 0 && newJ < len(board[0]) && !visited[newI][newJ] {
				if trace(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i := range board {
		for j := range board[i] {
			if trace(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end

