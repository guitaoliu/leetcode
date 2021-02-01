/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	maxReach := 0
	for i, v := range nums {
		if maxReach < i {
			return false
		}
		maxReach = max(maxReach, i+v)
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

