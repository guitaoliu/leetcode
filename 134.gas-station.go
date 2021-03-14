/*
 * @lc app=leetcode id=134 lang=golang
 *
 * [134] Gas Station
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	startAt := func(idx int) (int, bool) {
		cnt := 0
		for i := idx; i < idx+len(gas); i++ {
			pos := i % len(gas)
			cnt += gas[pos] - cost[pos]
			if cnt < 0 {
				return i, false
			}
		}
		return -1, true
	}
	for i := 0; i < len(gas); {
		if idx, canComplete := startAt(i); canComplete {
			return i
		} else {
			i = idx + 1
		}

	}
	return -1
}

// @lc code=end

