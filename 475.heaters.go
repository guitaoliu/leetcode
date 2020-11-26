/*
 * @lc app=leetcode id=475 lang=golang
 *
 * [475] Heaters
 */

// @lc code=start
func findRadius(houses []int, heaters []int) int {
	res := 0
	sort.Ints(heaters)
	for _, house := range houses {
		radius := findClosestHeater(house, heaters)
		if radius > res {
			res = radius
		}
	}
	return res
}

func findClosestHeater(house int, heaters []int) int {
	left, right := 0, len(heaters)-1
	for left < right {
		mid := left + (right-left)>>1
		if house < heaters[mid] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	if house == heaters[left] {
		return 0
	} else if house > heaters[left] {
		return house - heaters[left]
	} else if left == 0 {
		return heaters[left] - house
	} else {
		return min(heaters[left]-house, house-heaters[left-1])
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

