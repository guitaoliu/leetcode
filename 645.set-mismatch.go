/*
 * @lc app=leetcode id=645 lang=golang
 *
 * [645] Set Mismatch
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	m := make([]int, len(nums))
	for _, v := range nums {
		if m[v-1] == 0 {
			m[v-1] = 1
		} else {
			res[0] = v
		}
	}
	for idx, v := range m {
		if v == 0 {
			res[1] = idx + 1
			break
		}
	}
	return res
}

// @lc code=end

