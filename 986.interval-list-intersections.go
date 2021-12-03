/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */

// @lc code=start
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	for i, j := 0, 0; i < len(firstList) && j < len(secondList); {
		low := max(firstList[i][0], secondList[j][0])
		high := min(firstList[i][1], secondList[j][1])
		if low <= high {
			res = append(res, []int{low, high})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

