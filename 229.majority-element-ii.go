/*
 * @lc app=leetcode id=229 lang=golang
 *
 * [229] Majority Element II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	cnt1, cnt2 := 0, 0
	candidate1, candidate2 := 0, 1
	for _, v := range nums {
		if v == candidate1 {
			cnt1++
		} else if v == candidate2 {
			cnt2++
		} else if cnt1 == 0 {
			candidate1, cnt1 = v, 1
		} else if cnt2 == 0 {
			candidate2, cnt2 = v, 1
		} else {
			cnt1--
			cnt2--
		}
	}
	cnt1, cnt2 = 0, 0
	for _, v := range nums {
		if v == candidate1 {
			cnt1++
		} else if v == candidate2 {
			cnt2++
		}
	}
	res := make([]int, 0)
	if cnt1 > len(nums)/3 {
		res = append(res, candidate1)
	}
	if cnt2 > len(nums)/3 {
		res = append(res, candidate2)
	}
	return res
}

func majorityElementHash(nums []int) []int {
	m := make(map[int]int)
	boundary := len(nums) / 3
	res := []int{}
	for _, num := range nums {
		m[num]++
	}
	for num, v := range m {
		if v > boundary {
			res = append(res, num)
		}
	}
	return res
}

// @lc code=end

