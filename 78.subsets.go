/*
 * @lc app=leetcode id=78 lang=golang
 *
 * [78] Subsets
 */

// @lc code=start
// func subsets(nums []int) [][]int {
// 	res := make([][]int, 1)
// 	var backtracing func(int, []int)
// 	backtracing = func(idx int, result []int) {
// 		if len(result) > 0 {
// 			c := make([]int, len(result))
// 			copy(c, result)
// 			res = append(res, c)
// 		}
// 		for i := idx; i < len(nums); i++ {
// 			result = append(result, nums[i])
// 			backtracing(i+1, result)
// 			result = result[:len(result)-1]
// 		}
// 	}
// 	backtracing(0, []int{})
// 	return res
// }

func subsets(nums []int) [][]int {
	n := 1 << len(nums)
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		tmp := make([]int, 0)
		for j := 0; j < len(nums); j++ {
			if i>>j&1 == 1 {
				tmp = append(tmp, nums[j])
			}
		}
		res[i] = tmp
	}
	return res
}

// @lc code=end

