/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
func sortColors(nums []int) {
	cnt := make([]int, 3)
	for _, v := range nums {
		cnt[v]++
	}
	pos := 0
	for idx, num := range cnt {
		for k := 0; k < num; k++ {
			nums[pos] = idx
			pos++
		}
	}
}

// @lc code=end

