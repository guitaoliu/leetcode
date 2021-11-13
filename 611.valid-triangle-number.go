/*
 * @lc app=leetcode id=611 lang=golang
 *
 * [611] Valid Triangle Number
 */

// @lc code=start
func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res += sort.SearchInts(nums[j+1:], nums[i]+nums[j])
		}
	}
	return res
}

func triangleNumberTwoPointers(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := len(nums) - 1; i >= 0; i-- {
		l, r := 0, i-1
		for l < r {
			if nums[l]+nums[r] > nums[i] {
				res += r - l
				r--
			} else {
				l++
			}
		}
	}
	return res
}

// @lc code=end

