/*
 * @lc app=leetcode id=594 lang=golang
 *
 * [594] Longest Harmonious Subsequence
 */

// @lc code=start
func findLHS(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	numsMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numsMap[v]++
	}
	res := 0
	for k, v := range numsMap {
		if num, ok := numsMap[k+1]; ok {
			if num+v > res {
				res = num + v
			}
		}
	}
	return res
}

// @lc code=end

