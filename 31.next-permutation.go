/*
 * @lc app=leetcode id=31 lang=golang
 *
 * [31] Next Permutation
 */

// @lc code=start
func nextPermutation(nums []int) {
	l := len(nums) - 2
	for l >= 0 && nums[l] >= nums[l+1] {
		l--
	}
	if l >= 0 {
		r := len(nums) - 1
		for r >= 0 && nums[r] <= nums[l] {
			r--
		}
		nums[l], nums[r] = nums[r], nums[l]
	}
	reverse(nums[l+1:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

// @lc code=end

