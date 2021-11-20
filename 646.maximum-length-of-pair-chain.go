/*
 * @lc app=leetcode id=646 lang=golang
 *
 * [646] Maximum Length of Pair Chain
 */

// @lc code=start
func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	res := 0
	cur := math.MinInt64
	for _, pair := range pairs {
		if cur < pair[0] {
			cur = pair[1]
			res++
		}
	}
	return res
}
func findLongestChainDP(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })
	dp := make([]int, len(pairs))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	dp[0] = 1
	res := 0
	for i := 0; i < len(pairs); i++ {
		for j := 0; j < i; j++ {
			if pairs[j][1] < pairs[i][0] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

