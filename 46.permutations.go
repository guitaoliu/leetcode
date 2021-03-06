/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := make([][]int, 0)

	var backtrace func([]int, []bool, int)
	backtrace = func(result []int, used []bool, idx int) {
		if idx == len(nums) {
			tmp := make([]int, len(result))
			copy(tmp, result)
			res = append(res, tmp)
			return
		} else {
			for i := 0; i < len(nums); i++ {
				if !used[i] {
					used[i] = true
					result = append(result, nums[i])
					backtrace(result, used, idx+1)
					result = result[:len(result)-1]
					used[i] = false
				}
			}
		}
	}
	used := make([]bool, len(nums))
	backtrace([]int{}, used, 0)
	return res
}

// @lc code=end

