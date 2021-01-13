/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	start, end := 0, len(nums)-1
	length := len(nums)
	add := 0
	for idx := 1; idx < length-1; idx++ {
		start, end = 0, length-1
		if idx > 1 && nums[idx] == nums[idx-1] {
			start = idx - 1
		}
		for start < idx && end > idx {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			add = nums[start] + nums[end] + nums[idx]
			if add == 0 {
				res = append(res, []int{nums[start], nums[idx], nums[end]})
				start++
				end--
			} else if add > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return res
}

func threeSumOther(nums []int) [][]int {
	res := make([][]int, 0)
	counter := make(map[int]int, len(nums))
	uniqueNums := make([]int, 0)
	for _, v := range nums {
		counter[v]++
	}
	for idx := range counter {
		uniqueNums = append(uniqueNums, idx)
	}
	sort.Ints(uniqueNums)

	for i := 0; i < len(uniqueNums); i++ {
		if uniqueNums[i] == 0 && counter[uniqueNums[i]] >= 3 {
			res = append(res, []int{0, 0, 0})
		}
		for j := i + 1; j < len(uniqueNums); j++ {
			if uniqueNums[i]*2+uniqueNums[j] == 0 && counter[uniqueNums[i]] >= 2 {
				res = append(res, []int{uniqueNums[i], uniqueNums[i], uniqueNums[j]})
			}
			if uniqueNums[i]+uniqueNums[j]*2 == 0 && counter[uniqueNums[j]] >= 2 {
				res = append(res, []int{uniqueNums[i], uniqueNums[j], uniqueNums[j]})
			}

			c := 0 - uniqueNums[i] - uniqueNums[j]
			if _, ok := counter[c]; ok && c > uniqueNums[j] {
				res = append(res, []int{uniqueNums[i], uniqueNums[j], c})
			}
		}
	}
	return res
}

// @lc code=end

