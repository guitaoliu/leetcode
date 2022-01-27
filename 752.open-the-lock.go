/*
 * @lc app=leetcode id=752 lang=golang
 *
 * [752] Open the Lock
 */

// @lc code=start
func openLock(deadends []string, target string) int {
	if target == "0000" {
		return 0
	}
	targetNum, _ := strconv.Atoi(target)
	visited := make([]bool, 10000)
	for _, deadend := range deadends {
		num, _ := strconv.Atoi(deadend)
		if num == 0 {
			return -1
		}
		visited[num] = true
	}
	level := 0
	queue := []int{0}
	var next int
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		for _, num := range curLevel {
			for i := 1000; i > 0; i /= 10 {
				d := (num / i) % 10
				if d == 9 {
					next = num - 9*i
				} else {
					next = num + i
				}
				if next == targetNum {
					return level + 1
				}
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
				if d == 0 {
					next = num + 9*i
				} else {
					next = num - i
				}
				if next == targetNum {
					return level + 1
				}
				if !visited[next] {
					visited[next] = true
					queue = append(queue, next)
				}
			}
		}
		level++
	}
	return -1
}

// @lc code=end

