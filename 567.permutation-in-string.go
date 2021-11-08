/*
 * @lc app=leetcode id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	ch := make([]int, 26)
	for i := 0; i < m; i++ {
		ch[s1[i]-'a']--
	}
	left := 0
	for right := 0; right < n; right++ {
		ch[s2[right]-'a']++
		for ch[s2[right]-'a'] > 0 {
			ch[s2[left]-'a']--
			left++
		}
		if right-left+1 == m {
			return true
		}
	}
	return false
}

func checkInclusionSlideWindow(s1 string, s2 string) bool {
	m, n := len(s1), len(s2)
	if m > n {
		return false
	}
	ch1, ch2 := [26]int{}, [26]int{}
	for i := 0; i < m; i++ {
		ch1[s1[i]-'a']++
		ch2[s2[i]-'a']++
	}
	if ch1 == ch2 {
		return true
	}
	for i := m; i < n; i++ {
		ch2[s2[i]-'a']++
		ch2[s2[i-m]-'a']--
		if ch1 == ch2 {
			return true
		}
	}
	return false
}

// @lc code=end

