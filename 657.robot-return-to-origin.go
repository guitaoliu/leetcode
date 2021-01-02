/*
 * @lc app=leetcode id=657 lang=golang
 *
 * [657] Robot Return to Origin
 */

// @lc code=start
func judgeCircle(moves string) bool {
	m := make([]int, 26)
	for _, v := range moves {
		m[v-'A']++
	}
	return m['U'-'A'] == m['D'-'A'] &&
		m['L'-'A'] == m['R'-'A']
}

// @lc code=end

