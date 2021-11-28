/*
 * @lc app=leetcode id=120 lang=golang
 *
 * [120] Triangle
 */

// @lc code=start
func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i+1][j], triangle[i+1][j+1])
		}
	}
	return triangle[0][0]
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	res := math.MaxInt32
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else {
				triangle[i][j] += min(triangle[i-1][j], triangle[i-1][j-1])
			}
			if i == len(triangle)-1 {
				if triangle[i][j] < res {
					res = triangle[i][j]
				}
			}
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

