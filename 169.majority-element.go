/*
 * @lc app=leetcode id=169 lang=golang
 *
 * [169] Majority Element
 */

// @lc code=start
// func majorityElement(nums []int) int {
// 	m := make(map[int]int, 0)
// 	max, num := 0, 0
// 	for i := 0; i < len(nums); i++ {
// 		_, ok := m[nums[i]]
// 		if ok {
// 			m[nums[i]]++
// 		} else {
// 			m[nums[i]] = 1
// 		}
// 		if m[nums[i]] > max {
// 			max, num = m[nums[i]], nums[i]
// 		}
// 	}
// 	return num
// }

func majorityElement(nums []int) int {
	cnt, major := 1, nums[0]
	for i := 1; i < len(nums); i++ {
		if major != nums[i] {
			cnt--
		}
		if major == nums[i] || cnt == 0 {
			cnt++
			major = nums[i]
		}
	}
	return major
}

// @lc code=end

