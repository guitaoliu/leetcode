/*
 * @lc app=leetcode id=454 lang=golang
 *
 * [454] 4Sum II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	m := make(map[int]int, len(nums1)*len(nums2))
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			m[n1+n2]++
		}
	}
	cnt := 0
	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			cnt += m[-n3-n4]
		}
	}
	return cnt
}

// @lc code=end

