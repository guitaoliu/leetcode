/*
 * @lc app=leetcode id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]bool, len(nums1))
	for _, v := range nums1 {
		m[v] = true
	}
	for _, v := range nums2 {
		if m[v] {
			delete(m, v)
			res = append(res, v)
		}
	}
	return res
}

// @lc code=end

