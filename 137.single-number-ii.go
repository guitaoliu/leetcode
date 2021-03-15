/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	one, two := 0, 0
	for _, num := range nums {
		one = ^two & (one ^ num)
		two = ^one & (two ^ num)
	}
	return one
}

// @lc code=end

