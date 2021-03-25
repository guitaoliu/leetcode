/*
 * @lc app=leetcode id=152 lang=golang
 *
 * [152] Maximum Product Subarray
 */

// @lc code=start
func maxProduct(nums []int) int {
	maximun, minimun := nums[0], nums[0]
	res := nums[0]
	for _, num := range nums[1:] {
		maxM, minM := maximun, minimun
		maximun = max(num, num*maxM, num*minM)
		minimun = min(num, num*minM, num*maxM)
		res = max(maximun, res)
	}
	return res
}

func max(nums ...int) int {
	tmp := nums[0]
	for _, num := range nums[1:] {
		if num > tmp {
			tmp = num
		}
	}
	return tmp
}

func min(nums ...int) int {
	tmp := nums[0]
	for _, num := range nums[1:] {
		if num < tmp {
			tmp = num
		}
	}
	return tmp
}

// @lc code=end

