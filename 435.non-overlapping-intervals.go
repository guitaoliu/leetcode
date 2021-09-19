/*
 * @lc app=leetcode id=435 lang=golang
 *
 * [435] Non-overlapping Intervals
 */

// @lc code=start
import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			if intervals[i][1] <= intervals[i][1] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	})

	pre, cnt := 0, 1

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= intervals[pre][1] {
			cnt++
			pre = i
		} else if intervals[i][1] < intervals[pre][1] {
			pre = i
		}
	}
	return len(intervals) - cnt
}

func eraseOverlapIntervalsDP(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] {
			if intervals[i][1] <= intervals[i][1] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	})

	dp := make([]int, len(intervals))

	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(intervals); i++ {
		for j := 0; j < i; j++ {
			if intervals[i][0] >= intervals[j][1] {
				dp[i] = max(dp[i], 1+dp[j])
			}
		}
	}

	res := 0
	for _, v := range dp {
		res = max(res, v)
	}
	return len(intervals) - res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

