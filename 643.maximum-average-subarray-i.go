/*
 * @lc app=leetcode id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
	}
	res := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		res = max(sum, res)
	}
	return float64(res) / float64(k)

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

