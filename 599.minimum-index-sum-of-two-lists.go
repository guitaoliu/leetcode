/*
 * @lc app=leetcode id=599 lang=golang
 *
 * [599] Minimum Index Sum of Two Lists
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int, len(list1))
	res := []string{}
	for idx, v := range list1 {
		m[v] = idx
	}
	for idx, v := range list2 {
		if _, ok := m[v]; ok {
			m[v] += idx
			if len(res) == 0 || m[v] == m[res[0]] {
				res = append(res, v)
			} else if m[v] < m[res[0]] {
				res = []string{v}
			}
		}
	}
	return res
}

// @lc code=end

