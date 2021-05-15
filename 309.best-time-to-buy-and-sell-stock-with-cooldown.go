/*
 * @lc app=leetcode id=309 lang=golang
 *
 * [309] Best Time to Buy and Sell Stock with Cooldown
 */

// @lc code=start
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	p0, p1, p2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		p0New := max(p0, p2-prices[i])
		p1New := p0 + prices[i]
		p2New := max(p1, p2)
		p0, p1, p2 = p0New, p1New, p2New
	}
	return max(p1, p2)
}

func maxProfitDP(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	p0, p1, p2 := make([]int, n), make([]int, n), make([]int, n)
	p0[0] = -prices[0]
	for i := 1; i < n; i++ {
		p0[i] = max(p0[i-1], p2[i-1]-prices[i])
		p1[i] = p0[i-1] + prices[i]
		p2[i] = max(p1[i-1], p2[i-1])
	}
	return max(p1[n-1], p2[n-1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

