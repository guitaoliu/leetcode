/*
 * @lc app=leetcode id=274 lang=golang
 *
 * [274] H-Index
 */

// @lc code=start
func hIndex(citations []int) int {
	sort.Ints(citations)
	hIndex := 0
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= len(citations)-i {
			hIndex++
		} else {
			break
		}
	}
	return hIndex
}

// @lc code=end

