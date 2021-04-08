/*
 * @lc app=leetcode id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	left, right := 0, -1
	sum := 0
	res := math.MaxInt32
	for left < len(nums) {
		if right+1 < len(nums) && sum < target {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}
		if sum >= target {
			res = min(res, right-left+1)
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

