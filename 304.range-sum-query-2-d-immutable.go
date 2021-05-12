/*
 * @lc app=leetcode id=304 lang=golang
 *
 * [304] Range Sum Query 2D - Immutable
 */

// @lc code=start
type NumMatrix struct {
	sum [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 {
		return NumMatrix{nil}
	}
	sum := make([][]int, len(matrix)+1)
	sum[0] = make([]int, len(matrix[0])+1)
	for i := range matrix {
		sum[i+1] = make([]int, len(matrix[i])+1)
		for j := range matrix[i] {
			sum[i+1][j+1] = matrix[i][j] + sum[i][j+1] + sum[i+1][j] - sum[i][j]
		}
	}
	return NumMatrix{sum}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.sum[row2+1][col2+1] - this.sum[row1][col2+1] - this.sum[row2+1][col1] + this.sum[row1][col1]
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

