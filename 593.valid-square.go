/*
 * @lc app=leetcode id=593 lang=golang
 *
 * [593] Valid Square
 */

// @lc code=start
func validSquareSort(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	pointers := [][]int{p1, p2, p3, p4}
	sort.Slice(pointers, func(i, j int) bool {
		if pointers[i][0] < pointers[j][0] {
			return true
		} else if pointers[i][0] == pointers[j][0] {
			return pointers[i][1] < pointers[j][1]
		} else {
			return false
		}
	})
	return check(pointers[0], pointers[1], pointers[3], pointers[2])
}

func dis(p1, p2 []int) int {
	return (p1[0]-p2[0])*(p1[0]-p2[0]) + (p1[1]-p2[1])*(p1[1]-p2[1])
}

func check(p1, p2, p3, p4 []int) bool {
	return dis(p1, p2) > 0 && dis(p1, p2) == dis(p2, p3) && dis(p2, p3) == dis(p3, p4) && dis(p3, p4) == dis(p4, p1) && dis(p1, p3) == dis(p2, p4)
}

// @lc code=end

