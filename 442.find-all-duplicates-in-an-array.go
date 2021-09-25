/*
 * @lc app=leetcode id=442 lang=golang
 *
 * [442] Find All Duplicates in an Array
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	res := []int{}
	for i := 0; i < len(nums); i++ {
		cur := abs(nums[i])
		if nums[cur-1] < 0 {
			res = append(res, cur)
		} else {
			nums[cur-1] *= -1
		}

	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

