/*
 * @lc app=leetcode id=743 lang=golang
 *
 * [743] Network Delay Time
 */

// @lc code=start
func networkDelayTime(times [][]int, n int, k int) int {
	m := make(map[int][][2]int, n)
	for _, time := range times {
		m[time[0]] = append(m[time[0]], [2]int{time[1], time[2]})
	}
	reachTime := make([]int, n+1)
	for i := range reachTime {
		reachTime[i] = math.MaxInt32
	}
	queue := []int{k}
	reachTime[k] = 0
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		for _, node := range curLevel {
			nei := m[node]
			for _, nn := range nei {
				tt := reachTime[node] + nn[1]
				if tt >= reachTime[nn[0]] {
					continue
				}
				queue = append(queue, nn[0])
				reachTime[nn[0]] = tt
			}
		}
	}
	v := 0
	for i := 1; i <= n; i++ {
		if reachTime[i] == math.MaxInt32 {
			return -1
		}
		v = max(v, reachTime[i])
	}
	return v
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

