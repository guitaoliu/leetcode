/*
 * @lc app=leetcode id=557 lang=golang
 *
 * [557] Reverse Words in a String III
 */

// @lc code=start
func reverseWords(s string) string {
	sBytes := []byte(s)
	cur := 0
	for i := 0; i < len(s); i++ {
		if sBytes[i] == ' ' || i == len(s)-1 {
			j := i - 1
			if i == len(s)-1 {
				j++
			}
			for cur < j {
				sBytes[cur], sBytes[j] = sBytes[j], sBytes[cur]
				cur++
				j--
			}
			cur = i + 1
		}

	}
	return string(sBytes)
}

// @lc code=end

