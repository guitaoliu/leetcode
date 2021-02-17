/*
 * @lc app=leetcode id=80 lang=golang
 *
 * [80] Remove Duplicates from Sorted Array II
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	w, r := 1, 1
	cnt := 1
	for ; r < len(nums); r++ {
		if nums[r] == nums[r-1] {
			cnt++
		} else {
			cnt = 1
		}
		if cnt <= 2 {
			nums[w] = nums[r]
			w++
		}
	}
	return w
}

// @lc code=end

