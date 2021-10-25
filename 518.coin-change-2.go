/*
 * @lc app=leetcode id=518 lang=golang
 *
 * [518] Coin Change 2
 */

// @lc code=start
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}
	return dp[amount]
}

func changeDFS(amount int, coins []int) int {
	res := 0
	sort.Slice(coins, func(i, j int) bool { return coins[i] > coins[j] })
	var trace func(int, int)
	trace = func(remain int, index int) {
		if remain == 0 {
			res++
			return
		} else {
			for i := index; i < len(coins); i++ {
				if coins[i] <= remain {
					trace(remain-coins[i], i)
				}
			}
		}
	}
	trace(amount, 0)
	return res
}

// @lc code=end

