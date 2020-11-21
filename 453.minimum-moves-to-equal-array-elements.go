/*
 * @lc app=leetcode id=453 lang=golang
 *
 * [453] Minimum Moves to Equal Array Elements
 */

// @lc code=start
func minMoves(nums []int) int {
	min, sum := math.MaxInt32, 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if nums[i] < min {
			min = nums[i]
		}
	}
	return sum - len(nums)*min
}

// @lc code=end

 