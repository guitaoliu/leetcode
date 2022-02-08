/*
 * @lc app=leetcode id=754 lang=golang
 *
 * [754] Reach a Number
 */

// @lc code=start
func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum := 0
	k := 0
	for sum < target {
		sum += (k + 1)
		k++
	}
	d := sum - target
	if d%2 == 0 {
		return k
	}
	return k + 1 + k%2
}

// @lc code=end

