/*
 * @lc app=leetcode id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func arrangeCoins(n int) int {
	k := 1
	for k*(1+k)/2 <= n {
		k++
	}
	k--
	return k
}

// @lc code=end

