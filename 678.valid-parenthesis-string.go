/*
 * @lc app=leetcode id=678 lang=golang
 *
 * [678] Valid Parenthesis String
 */

// @lc code=start
func checkValidString(s string) bool {
	leftPStack := []int{}
	starStack := []int{}
	for i, ch := range s {
		switch ch {
		case '(':
			leftPStack = append(leftPStack, i)
		case '*':
			starStack = append(starStack, i)
		case ')':
			if len(leftPStack) > 0 {
				leftPStack = leftPStack[:len(leftPStack)-1]
			} else if len(starStack) > 0 {
				starStack = starStack[:len(starStack)-1]
			} else {
				return false
			}
		}
	}
	i, j := len(leftPStack)-1, len(starStack)-1
	for ; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if leftPStack[i] > starStack[j] {
			return false
		}
	}
	return i < 0
}

// @lc code=end

