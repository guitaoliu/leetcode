/*
 * @lc app=leetcode id=844 lang=golang
 *
 * [844] Backspace String Compare
 */

// @lc code=start
func backspaceCompare(s string, t string) bool {
	stack1 := []byte{}
	stack2 := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == '#' {
			if len(stack1) > 0 {
				stack1 = stack1[:len(stack1)-1]
			}
		} else {
			stack1 = append(stack1, s[i])
		}
	}
	for i := 0; i < len(t); i++ {
		if t[i] == '#' {
			if len(stack2) > 0 {
				stack2 = stack2[:len(stack2)-1]
			}
		} else {
			stack2 = append(stack2, t[i])
		}
	}
	if len(stack1) != len(stack2) {
		return false
	}
	for i := 0; i < len(stack1); i++ {
		if stack1[i] != stack2[i] {
			return false
		}
	}
	return true
}

// @lc code=end

