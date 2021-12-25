/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
    i := 0
    for i < len(nums) {
        if nums[i] < len(nums) && nums[i] != i {
            j := nums[i]
            nums[i], nums[j] = nums[j], nums[i]
        } else {
            i++
        }
    }
    for i, n := range nums {
        if i != n {
            return i
        }
    }
    return len(nums)
}

func missingNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		res = res + i + 1 - nums[i]
	}
	return res
}

// @lc code=end

