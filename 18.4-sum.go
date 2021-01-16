/*
 * @lc app=leetcode id=18 lang=golang
 *
 * [18] 4Sum
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left, right := j+1, len(nums)-1
			for left < right {
				if left > j+1 && nums[left] == nums[left-1] {
					left++
					continue
				}
				if right < len(nums)-1 && nums[right] == nums[right+1] {
					right--
					continue
				}
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					left++
					right--
				} else if sum > target {
					right--
				} else {
					left++
				}
			}
		}

	}
	return res
}

// @lc code=end

