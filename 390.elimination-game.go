/*
 * @lc app=leetcode id=390 lang=golang
 *
 * [390] Elimination Game
 */

// @lc code=start
func lastRemaining(n int) int {
	head, step := 1, 1
	left := true
	remain := n
	for remain > 1 {
		if left || remain%2 != 0 {
			head += step
		}

		remain /= 2
		step *= 2
		left = !left
	}
	return head
}

// @lc code=end

