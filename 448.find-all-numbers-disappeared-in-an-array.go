/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	for _, v := range nums {
		if v < 0 {
			v = -v
		}
		if nums[v-1] > 0 {
			nums[v-1] = -nums[v-1]
		}
	}
	res := make([]int, 0)
	for idx, v := range nums {
		if v > 0 {
			res = append(res, idx+1)
		}
	}
	return res
}

// @lc code=end

