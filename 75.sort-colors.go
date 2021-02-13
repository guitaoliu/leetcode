/*
 * @lc app=leetcode id=75 lang=golang
 *
 * [75] Sort Colors
 */

// @lc code=start
// func sortColors(nums []int) {
// 	cnt := make([]int, 3)
// 	for _, v := range nums {
// 		cnt[v]++
// 	}
// 	pos := 0
// 	for idx, num := range cnt {
// 		for k := 0; k < num; k++ {
// 			nums[pos] = idx
// 			pos++
// 		}
// 	}
// }

func sortColors(nums []int) {
	p0, p1 := 0, 0
	for idx, color := range nums {
		if color == 0 {
			nums[idx], nums[p0] = nums[p0], nums[idx]
			if p0 < p1 {
				nums[idx], nums[p1] = nums[p1], nums[idx]
			}
			p1++
			p0++
		} else if color == 1 {
			nums[idx], nums[p1] = nums[p1], nums[idx]
			p1++
		}
	}
}

// @lc code=end

