/*
 * @lc app=leetcode id=46 lang=golang
 *
 * [46] Permutations
 */

// @lc code=start
func permute(nums []int) [][]int {
	res := make([][]int, 0)

	var backtrace func([]int, int, map[int]struct{})
	backtrace = func(result []int, idx int, used map[int]struct{}) {
		if idx == len(nums) {
			c := make([]int, len(result))
			copy(c, result)
			res = append(res, c)
		} else {
			for _, v := range nums {
				if _, ok := used[v]; !ok {
					used[v] = struct{}{}
					backtrace(append(result, v), idx+1, used)
					delete(used, v)
				}
			}
		}
	}
	used := make(map[int]struct{}, 0)
	backtrace([]int{}, 0, used)
	return res
}

// @lc code=end

