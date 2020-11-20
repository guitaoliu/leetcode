/*
 * @lc app=leetcode id=447 lang=golang
 *
 * [447] Number of Boomerangs
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	res := 0
	for i := 0; i < len(points); i++ {
		m := make(map[int]int, len(points))
		for j := 0; j < len(points); j++ {
			if i != j {
				m[dis(points[i], points[j])]++
			}
		}
		for _, v := range m {
			res += v * (v - 1)
		}
	}
	return res
}

func dis(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

// @lc code=end

