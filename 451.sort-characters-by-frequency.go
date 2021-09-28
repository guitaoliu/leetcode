/*
 * @lc app=leetcode id=451 lang=golang
 *
 * [451] Sort Characters By Frequency
 */

// @lc code=start
import (
	"strings"
)

func frequencySort(s string) string {
	freq := make(map[byte]int)
	max := 0
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
		if freq[s[i]] > max {
			max = freq[s[i]]
		}
	}
	mm := make(map[int][]byte)
	for k, v := range freq {
		mm[v] = append(mm[v], k)
	}

	res := make([]byte, 0, len(s))
	for max >= 0 {
		if v, ok := mm[max]; ok {
			for _, ch := range v {
				for i := 0; i < max; i++ {
					res = append(res, ch)
				}
			}
		}
		max--
	}
	return string(res)
}

// @lc code=end

