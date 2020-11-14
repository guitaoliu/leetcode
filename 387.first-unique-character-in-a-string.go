/*
 * @lc app=leetcode id=387 lang=golang
 *
 * [387] First Unique Character in a String
 */

// @lc code=start
func firstUniqChar(s string) int {
	res := make([]int, 26)
	for _, v := range s {
		res[v-'a']++
	}
	for idx, v := range s {
		if res[v-'a'] == 1 {
			return idx
		}
	}
	return -1
}

// @lc code=end

