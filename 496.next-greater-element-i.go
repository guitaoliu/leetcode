/*
 * @lc app=leetcode id=496 lang=golang
 *
 * [496] Next Greater Element I
 */

// @lc code=start
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	res := make([]int, 0)
	n1 := make(map[int]int, len(nums2))
	stack := make([]int, 0)
	for i := 0; i < len(nums2); i++ {
		for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
			n1[stack[len(stack)-1]] = nums2[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}
	for _, v := range nums1 {
		if idx, ok := n1[v]; ok {
			res = append(res, idx)
		} else {
			res = append(res, -1)
		}
	}
	return res
}

// @lc code=end

