/*
 * @lc app=leetcode id=556 lang=golang
 *
 * [556] Next Greater Element III
 */

// @lc code=start

func nextGreaterElement(n int) int {
	ss := strconv.Itoa(n)
	s := []byte(ss)
	i := len(s) - 2
	for i >= 0 && s[i] >= s[i+1] {
		i--
	}
	if i < 0 {
		return -1
	}
	j := len(s) - 1
	for j >= 0 && s[j] <= s[i] {
		j--
	}
	s[i], s[j] = s[j], s[i]
	for m, n := i+1, len(s)-1; m < n; m, n = m+1, n-1 {
		s[m], s[n] = s[n], s[m]
	}
	res, _ := strconv.Atoi(string(s))
	if res > math.MaxInt32 {
		return -1
	}
	return res
}

// @lc code=end

