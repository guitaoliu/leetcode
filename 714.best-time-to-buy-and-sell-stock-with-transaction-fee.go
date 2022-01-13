/*
 * @lc app=leetcode id=714 lang=golang
 *
 * [714] Best Time to Buy and Sell Stock with Transaction Fee
 */

// @lc code=start
func maxProfit(prices []int, fee int) int {
	buy := prices[0] + fee
	profit := 0
	for _, price := range prices[1:] {
		if price+fee < buy {
			buy = price + fee
		} else if price > buy {
			profit += price - buy
			buy = price
		}
	}
	return profit
}

func maxProfitDP(prices []int, fee int) int {
	if len(prices) <= 1 {
		return 0
	}
	buy := make([]int, len(prices))
	sell := make([]int, len(prices))
	buy[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		buy[i] = max(buy[i-1], sell[i-1]-prices[i])
		sell[i] = max(sell[i-1], buy[i-1]+prices[i]-fee)
	}
	return sell[len(sell)-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

