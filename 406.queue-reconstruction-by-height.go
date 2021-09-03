/*
 * @lc app=leetcode id=406 lang=golang
 *
 * [406] Queue Reconstruction by Height
 */

// @lc code=start
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] < people[j][0] || people[i][0] == people[j][0] && people[i][1] > people[j][1]
	})
	res := make([][]int, len(people))
	for _, p := range people {
		cnt := p[1] + 1
		for i := range res {
			if res[i] == nil {
				cnt--
				if cnt == 0 {
					res[i] = p
				}
			}
		}
	}
	return res
}

func reconstructQueueHigh2Low(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
	res := [][]int{}
	for _, p := range people {
		idx := p[1]
		res = append(res[:idx], append([][]int{p}, res[idx:]...)...)
	}
	return res
}

// @lc code=end

