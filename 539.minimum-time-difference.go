/*
 * @lc app=leetcode id=539 lang=golang
 *
 * [539] Minimum Time Difference
 */

// @lc code=start
func findMinDifference(timePoints []string) int {
	minutes := make([]int, len(timePoints))
	for i, time := range timePoints {
		h, _ := strconv.Atoi(time[:2])
		m, _ := strconv.Atoi(time[3:])
		minutes[i] = h*60 + m
	}
	sort.Ints(minutes)
	res := math.MaxInt64
	for i := 0; i < len(timePoints)-1; i++ {
		res = min(res, minutes[i+1]-minutes[i])
	}
	res = min(res, 24*60-(minutes[len(minutes)-1]-minutes[0]))
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

