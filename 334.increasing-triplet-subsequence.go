/*
 * @lc app=leetcode id=334 lang=golang
 *
 * [334] Increasing Triplet Subsequence
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	i, j := math.MaxInt32, math.MaxInt32
	for _, num := range nums {
		if num > j {
			return true
		} else if num > i {
			j = num
		} else {
			i = num
		}
	}
	return false
}

// @lc code=end

