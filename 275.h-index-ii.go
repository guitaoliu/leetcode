/*
 * @lc app=leetcode id=275 lang=golang
 *
 * [275] H-Index II
 */

// @lc code=start
func hIndex(citations []int) int {
	left, right := 0, len(citations)-1
	for left <= right {
		mid := left + (right-left)>>1
		if citations[mid] < len(citations)-mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return len(citations) - left
}

// @lc code=end

