/*
 * @lc app=leetcode id=16 lang=golang
 *
 * [16] 3Sum Closest
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res, maxDiff := 0, math.MaxInt64
	sum := 0
	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum = nums[i] + nums[j] + nums[k]
			if diff := abs(sum - target); diff < maxDiff {
				res, maxDiff = sum, diff
			}
			if sum == target {
				return res
			} else if sum > target {
				k--
			} else {
				j++
			}
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

