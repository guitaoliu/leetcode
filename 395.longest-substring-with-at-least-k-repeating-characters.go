/*
 * @lc app=leetcode id=395 lang=golang
 *
 * [395] Longest Substring with At Least K Repeating Characters
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	if s == "" {
		return 0
	}

	if len(s) < k {
		return 0
	}

	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	var split byte
	for i, c := range cnt {
		if 0 < c && c < k {
			split = 'a' + byte(i)
			break
		}
	}

	if split == 0 {
		return len(s)
	}

	res := 0
	strings := strings.Split(s, string(split))
	for _, subString := range strings {
		res = max(res, longestSubstring(subString, k))
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

