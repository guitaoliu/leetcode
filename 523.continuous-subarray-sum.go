/*
 * @lc app=leetcode id=523 lang=golang
 *
 * [523] Continuous Subarray Sum
 */

// @lc code=start
func checkSubarraySumPrefix(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := make([]int, len(nums)+1)
	sum[0] = 0
	for i, num := range nums {
		sum[i+1] = sum[i] + num
		for j := i - 1; j >= 0; j-- {
			if (sum[i+1]-sum[j])%k == 0 {
				return true
			}
		}
	}
	return false
}

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	m := make(map[int]int)
	m[0] = -1
	sum := 0
	for i, num := range nums {
		sum += num
		if r, ok := m[sum%k]; ok {
			if i-2 >= r {
				return true
			}
		} else {
			m[sum%k] = i
		}
	}
	return false
}

// @lc code=end

