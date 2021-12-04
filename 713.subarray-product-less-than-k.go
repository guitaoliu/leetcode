/*
 * @lc app=leetcode id=713 lang=golang
 *
 * [713] Subarray Product Less Than K
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	i := 0
	res := 0
	prod := 1
	for j := 0; j < len(nums); j++ {
		if prod *= nums[j]; prod >= k {
			for i < len(nums) && prod >= k {
				prod /= nums[i]
				i++
			}
		}
		if prod < k {
			res += j - i + 1
		}
	}
	return res
}

// @lc code=end

