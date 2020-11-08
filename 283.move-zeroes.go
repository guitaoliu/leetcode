/*
 * @lc app=leetcode id=283 lang=golang
 *
 * [283] Move Zeroes
 */

// @lc code=start
func moveZeroes(nums []int) {
	head := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			if head != i {
				nums[head], nums[i] = nums[i], nums[head]
			}
			head++
		}
	}
}

// @lc code=end

