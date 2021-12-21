/*
 * @lc app=leetcode id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 */

// @lc code=start
func totalFruit(fruits []int) int {
	m := make(map[int]int)
	left, right := 0, 0
	res := 0
	for right < len(fruits) {
		m[fruits[right]]++
		for len(m) > 2 {
			m[fruits[left]]--
			if m[fruits[left]] == 0 {
				delete(m, fruits[left])
			}
			left++
		}
		if l := right - left + 1; l > res {
			res = l
		}
		right++
	}
	return res
}

// @lc code=end

