/*
 * @lc app=leetcode id=492 lang=golang
 *
 * [492] Construct the Rectangle
 */

// @lc code=start
func constructRectangle(area int) []int {
	for i := sqrtInt(area); i >= 1; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{area, 1}
}

func sqrtInt(n int) int {
	return int(math.Sqrt(float64(n)))
}

// @lc code=end

