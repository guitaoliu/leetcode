/*
 * @lc app=leetcode id=560 lang=golang
 *
 * [560] Subarray Sum Equals K
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	m := map[int]int{0: 1}
	res := 0
	sum := 0
	for _, num := range nums {
		sum += num
		if v, ok := m[sum-k]; ok {
			res += v
		}
		m[sum]++
	}
	return res
}

// @lc code=end

