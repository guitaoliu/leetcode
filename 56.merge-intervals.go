/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	res := make([][]int, 0)
	tmp := intervals[0]
	for _, v := range intervals[1:] {
		if v[0] > tmp[1] {
			res = append(res, tmp)
			tmp = v
		} else {
			tmp[1] = max(tmp[1], v[1])
		}
	}
	res = append(res, tmp)
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// @lc code=end

