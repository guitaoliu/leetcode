/*
 * @lc app=leetcode id=503 lang=golang
 *
 * [503] Next Greater Element II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}

	stack := []int{}
	for i := 0; i < len(nums)*2; i++ {
		num := nums[i%len(nums)]
		for len(stack) > 0 && nums[stack[len(stack)-1]] < num {
			top := stack[len(stack)-1]
			res[top] = num
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%len(nums))
	}

	return res
}

// @lc code=end

