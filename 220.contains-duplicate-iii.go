/*
 * @lc app=leetcode id=220 lang=golang
 *
 * [220] Contains Duplicate III
 */

// @lc code=start
func containsnearbyalmostduplicate(nums []int, k int, t int) bool {
	buckets := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		key := nums[i] / (t + 1)
		if nums[i] < 0 {
			key--
		}
		if _, ok := buckets[key]; ok {
			return true
		}
		if v, ok := buckets[key-1]; ok && nums[i]-v <= t {
			return true
		}
		if v, ok := buckets[key+1]; ok && v-nums[i] <= t {
			return true
		}
		if len(buckets) >= k {
			delete(buckets, nums[i-k]/(t+1))
		}
		buckets[key] = nums[i]
	}
	return false
}

// func containsnearbyalmostduplicate(nums []int, k int, t int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j <= i+k && j < len(nums); j++ {
// 			if abs(nums[i]-nums[j]) <= t {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// func abs(a int) int {
// 	if a > 0 {
// 		return a
// 	}
// 	return -a
// }

// @lc code=end

