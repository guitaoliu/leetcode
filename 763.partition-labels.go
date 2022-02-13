/*
 * @lc app=leetcode id=763 lang=golang
 *
 * [763] Partition Labels
 */

// @lc code=start
func partitionLabels(s string) []int {
	m := make(map[byte]int)
	sBytes := []byte(s)
	for _, c := range sBytes {
		m[c]++
	}
	pos := []int{}
	i := 0
	cur := make(map[byte]bool)
	for i < len(s) {
		if len(cur) == 0 {
			pos = append(pos, i)
		}
		if _, ok := cur[sBytes[i]]; !ok {
			cur[sBytes[i]] = true
		}
		m[sBytes[i]]--
		if m[sBytes[i]] == 0 {
			delete(cur, sBytes[i])
		}
		i++
	}
	res := make([]int, len(pos))
	for i := range pos[:len(pos)-1] {
		res[i] = pos[i+1] - pos[i]
	}
	res[len(res)-1] = len(sBytes) - pos[len(pos)-1]
	return res
}

// @lc code=end

