/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	res := make([][]int, 0, len(intervals))
	merged := false
	for _, v := range intervals {
		if v[0] > newInterval[1] {
			if !merged {
				res = append(res, newInterval)
				merged = true
			}
			res = append(res, v)
		} else if v[1] < newInterval[0] {
			res = append(res, v)
		} else {
			newInterval = extend(newInterval, v)
		}
	}
	if !merged {
		res = append(res, newInterval)
	}
	return res
}

func extend(a, b []int) []int {
	res := make([]int, 2)
	if a[0] < b[0] {
		res[0] = a[0]
	} else {
		res[0] = b[0]
	}
	if a[1] > b[1] {
		res[1] = a[1]
	} else {
		res[1] = b[1]
	}
	return res
}

// @lc code=end

