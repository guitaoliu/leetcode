/*
 * @lc app=leetcode id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	head1, head2 := 0, 0
	res := []int{}
	for head1 < len(nums1) && head2 < len(nums2) {
		if nums1[head1] < nums2[head2] {
			head1++
		} else if nums1[head1] > nums2[head2] {
			head2++
		} else {
			res = append(res, nums1[head1])
			head1++
			head2++
		}
	}
	return res
}

// @lc code=end

