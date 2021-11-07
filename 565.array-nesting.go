/*
 * @lc app=leetcode id=565 lang=golang
 *
 * [565] Array Nesting
 */

// @lc code=start
func arrayNestingMemo(nums []int) int {
	visited := make([]bool, len(nums))
	res := 0
	for i := range nums {
		cnt := 0
		idx := i
		for !visited[idx] {
			visited[idx] = true
			idx = nums[idx]
			cnt++
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}

func arrayNesting(nums []int) int {
	res := 0
	for i := range nums {
		cnt := 0
		idx := i
		for nums[idx] >= 0 {
			cnt++
			tmp := nums[idx]
			nums[idx] = -1
			idx = tmp
		}
		if cnt > res {
			res = cnt
		}
	}
	return res
}

// @lc code=end

