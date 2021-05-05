/*
 * @lc app=leetcode id=279 lang=golang
 *
 * [279] Perfect Squares
 */

// @lc code=start

func numSquares(n int) int {
	queue := []int{n}
	visited := make(map[int]struct{})
	visited[n] = struct{}{}
	level := 0
	for len(queue) > 0 {
		level++
		curLevel := queue
		queue = nil
		for _, node := range curLevel {
			for i := 1; i*i <= node; i++ {
				remain := node - i*i
				if remain == 0 {
					return level
				}
				if _, ok := visited[remain]; !ok {
					queue = append(queue, remain)
					visited[remain] = struct{}{}
				}
			}
		}
	}
	return level
}

// func numSquares(n int) int {
// 	dp := make([]int, n+1)
// 	dp[0] = 0
// 	for i := 1; i <= n; i++ {
// 		dp[i] = i
// 		for j := 1; i >= j*j; j++ {
// 			dp[i] = min(dp[i], dp[i-j*j]+1)
// 		}
// 	}
// 	return dp[n]
// }

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// @lc code=end

