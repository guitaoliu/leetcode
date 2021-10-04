/*
 * @lc app=leetcode id=456 lang=golang
 *
 * [456] 132 Pattern
 */

// @lc code=start
func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	numi := nums[0]
	for j := 1; j < len(nums); j++ {
		for k := len(nums) - 1; k > j; k-- {
			if numi < nums[k] && nums[k] < nums[j] {
				return true
			}
		}
		if nums[j] < numi {
			numi = nums[j]
		}
	}
	return false
}

func find132patternStack(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	stack := []int{}
	numk := math.MinInt32
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < numk {
			return true
		}
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			numk = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}

	return false
}

// @lc code=end

