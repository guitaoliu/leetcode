/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
// func maxSubArray(nums []int) int {
// 	if len(nums) == 0 {
// 		return 0
// 	}
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}
// 	maxToI, res := make([]int, len(nums)), nums[0]
// 	maxToI[0] = nums[0]
// 	for i := 1; i < len(nums); i++ {
// 		if maxToI[i-1] > 0 {
// 			maxToI[i] = maxToI[i-1] + nums[i]
// 		} else {
// 			maxToI[i] = nums[i]
// 		}
// 		if maxToI[i] > res {
// 			res = maxToI[i]
// 		}
// 	}
// 	return res
// }

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	maxSum, res, p := nums[0], 0, 0
	for p < len(nums) {
		res += nums[p]
		if res > maxSum {
			maxSum = res
		}
		if res < 0 {
			res = 0
		}
		p++
	}
	return maxSum
}

// @lc code=end

