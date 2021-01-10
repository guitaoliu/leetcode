/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		width := right - left
		var h int
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		if cur := h * width; cur > max {
			max = cur
		}
	}
	return max
}

// @lc code=end

