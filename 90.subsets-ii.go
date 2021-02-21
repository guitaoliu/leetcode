/*
 * @lc app=leetcode id=90 lang=golang
 *
 * [90] Subsets II
 */

// @lc code=start
func subsetsWithDup(nums []int) [][]int {
	res := make([][]int, 1)
	sort.Ints(nums)
	var backtracing func(int, []int)
	backtracing = func(idx int, result []int) {
		if len(result) > 0 {
			c := make([]int, len(result))
			copy(c, result)
			res = append(res, c)
		}
		for i := idx; i < len(nums); i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			result = append(result, nums[i])
			backtracing(i+1, result)
			result = result[:len(result)-1]
		}
	}
	backtracing(0, []int{})
	return res
}

// @lc code=end

