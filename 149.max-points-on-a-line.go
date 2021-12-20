/*
 * @lc app=leetcode id=149 lang=golang
 *
 * [149] Max Points on a Line
 */

// @lc code=start
func maxPoints(points [][]int) int {
	if len(points) <= 2 {
		return len(points)
	}
	ans := 0
	for i, p := range points {
		if ans > len(points)/2 || ans >= len(points)-i {
			break
		}
		m := make(map[int]int)
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			m[y+x<<16]++
		}
		for _, v := range m {
			ans = max(ans, v+1)
		}
	}
	return ans
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

