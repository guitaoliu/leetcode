/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	sort.Ints(candidates)
	var backtracing func(int, []int, int)
	backtracing = func(idx int, result []int, remaining int) {
		if remaining == 0 {
			c := make([]int, len(result))
			copy(c, result)
			res = append(res, c)
		} else {
			for i := idx; i < len(candidates); i++ {
				if i > idx && candidates[i] == candidates[i-1] {
					continue
				}
				if candidates[i] > target {
					return
				}
				backtracing(i+1, append(result, candidates[i]), remaining-candidates[i])
			}
		}
	}
	backtracing(0, []int{}, target)
	return res
}

// @lc code=end

