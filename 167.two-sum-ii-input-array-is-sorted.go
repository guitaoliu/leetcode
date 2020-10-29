/*
 * @lc app=leetcode id=167 lang=golang
 *
 * [167] Two Sum II - Input array is sorted
 */

// @lc code=start
func twoSum(numbers []int, target int) []int {
	num := make(map[int]int, 0)
	for i := 0; i < len(numbers); i++ {
		res := target - numbers[i]
		if _, ok := num[res]; ok {
			return []int{num[res] + 1, i + 1}
		} else {
			num[numbers[i]] = i
		}
	}
	return nil
}

// @lc code=end

