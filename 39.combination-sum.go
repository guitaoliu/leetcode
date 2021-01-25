/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	sort.Ints(candidates)
	var backtracing func(int, []int, int)
	backtracing = func(idx int, result []int, remaining int) {
		if remaining == 0 {
			c := make([]int, len(result))
			copy(c, result)
			res = append(res, c)
		} else {
			for i, v := range candidates[idx:] {
				if v > remaining {
					return
				}
				backtracing(idx+i, append(result, v), remaining-v)
			}
		}
	}
	backtracing(0, []int{}, target)
	return res
}

// @lc code=end

