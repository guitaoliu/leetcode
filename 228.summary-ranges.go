/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	res := []string{}
	for i := 0; i < len(nums); {
		start := i
		for i++; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		}
		s := strconv.Itoa(nums[start])
		if start != i-1 {
			s += "->" + strconv.Itoa(nums[i-1])
		}
		res = append(res, s)
	}
	return res
}

// @lc code=end

