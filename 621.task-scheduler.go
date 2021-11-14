/*
 * @lc app=leetcode id=621 lang=golang
 *
 * [621] Task Scheduler
 */

// @lc code=start
func leastInterval(tasks []byte, n int) int {
	cnt := make(map[byte]int)
	for _, t := range tasks {
		cnt[t]++
	}
	maxExec, maxExecCnt := 0, 0
	for _, c := range cnt {
		if c > maxExec {
			maxExec = c
			maxExecCnt = 1
		} else if c == maxExec {
			maxExecCnt++
		}
	}

	time := (maxExec-1)*(n+1) + maxExecCnt
	if len(tasks) > time {
		return len(tasks)
	}
	return time
}

func leastIntervalSimulate(tasks []byte, n int) int {
	cnt := make(map[byte]int)
	for _, task := range tasks {
		cnt[task]++
	}
	remainWorks := make([]int, 0, len(cnt))
	nextExecTime := make([]int, 0, len(cnt))
	for _, c := range cnt {
		nextExecTime = append(nextExecTime, 1)
		remainWorks = append(remainWorks, c)
	}

	currentTime := 0
	for i := 0; i < len(tasks); i++ {
		currentTime++
		// find the earliest start time with available works
		earliestStart := math.MaxInt64
		for idx, r := range remainWorks {
			if r > 0 && nextExecTime[idx] < earliestStart {
				earliestStart = nextExecTime[idx]
			}
		}

		if earliestStart > currentTime {
			currentTime = earliestStart
		}

		// find the one could start at earliestStart but has most remaining work
		execIndex := -1
		for idx, r := range remainWorks {
			if r > 0 && nextExecTime[idx] == earliestStart && (execIndex == -1 || r > remainWorks[execIndex]) {
				execIndex = idx
			}
		}

		// execute and set the next start time of that task
		nextExecTime[execIndex] = currentTime + n + 1
		remainWorks[execIndex]--
	}
	return currentTime
}

// @lc code=end

