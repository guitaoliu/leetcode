/*
 * @lc app=leetcode id=77 lang=golang
 *
 * [77] Combinations
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	var backtracing func(int, int, []int)
	backtracing = func(idx int, val int, result []int) {
		if idx == k {
			c := make([]int, len(result))
			copy(c, result)
			res = append(res, c)
		} else {
			for i := val; i <= n; i++ {
				if len(result)+n-i+1 >= k {
					backtracing(idx+1, i+1, append(result, i))
				}
			}
		}
	}
	backtracing(0, 1, []int{})
	return res
}

// @lc code=end

