/*
 * @lc app=leetcode id=436 lang=golang
 *
 * [436] Find Right Interval
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	kv := make([][]int, len(intervals))
	for i, interval := range intervals {
		kv[i] = []int{interval[0], i}
	}
	sort.Slice(kv, func(i, j int) bool { return kv[i][0] < kv[j][0] })

	find := func(target int) int {
		left, right := 0, len(kv)-1
		for left < right {
			mid := left + (right-left)>>1
			if kv[mid][0] == target {
				return kv[mid][1]
			} else if kv[mid][0] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if kv[left][0] >= target {
			return kv[left][1]
		}
		return -1
	}

	ans := []int{}
	for _, interval := range intervals {
		loc := find(interval[1])
		ans = append(ans, loc)
	}

	return ans
}

// @lc code=end

