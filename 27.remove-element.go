/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != idx {
				nums[i], nums[idx] = nums[idx], nums[i]
			}
			idx++
		}
	}
	return idx
}

// @lc code=end

