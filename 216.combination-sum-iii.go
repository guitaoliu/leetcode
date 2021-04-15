/*
 * @lc app=leetcode id=216 lang=golang
 *
 * [216] Combination Sum III
 */

// @lc code=start

func combinationSum3(k int, n int) [][]int {
	if k == 0 {
		return nil
	}
	res := make([][]int, 0)
	var backtracing func(int, int, []int)
	backtracing = func(remaining, idx int, result []int) {
		if len(result) > k {
			return
		}
		if remaining == 0 && len(result) == k {
			c := make([]int, len(result))
			copy(c, result)
			res = append(res, c)
		} else {
			for i := idx; i < 10; i++ {
				if i > remaining {
					break
				}
				result = append(result, i)
				backtracing(remaining-i, i+1, result)
				result = result[:len(result)-1]
			}
		}
	}
	backtracing(n, 1, []int{})
	return res
}

// @lc code=end

