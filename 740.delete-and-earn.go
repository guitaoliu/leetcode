/*
 * @lc app=leetcode id=740 lang=golang
 *
 * [740] Delete and Earn
 */

// @lc code=start
func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	maxVal := 0
	for _, num := range nums {
		maxVal = max(num, maxVal)
	}
	cnt := make([]int, maxVal+1)
	for _, num := range nums {
		cnt[num] += num
	}
	a, b := cnt[0], max(cnt[0], cnt[1])
	for i := 2; i < len(cnt); i++ {
		a, b = b, max(a+cnt[i], b)
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

