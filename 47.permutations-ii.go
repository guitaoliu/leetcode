/*
 * @lc app=leetcode id=47 lang=golang
 *
 * [47] Permutations II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	res := make([][]int, 0)
	sort.Ints(nums)
	var backtrace func([]int, []bool, int)
	backtrace = func(result []int, used []bool, idx int) {
		if idx == len(nums) {
			tmp := make([]int, len(result))
			copy(tmp, result)
			res = append(res, tmp)
		} else {
			for i := 0; i < len(nums); i++ {
				if !used[i] {
					if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
						continue
					}
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

