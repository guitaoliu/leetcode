/*
 * @lc app=leetcode id=227 lang=golang
 *
 * [227] Basic Calculator II
 */

// @lc code=start
func calculate(s string) int {
	stack := make([]int, 0)
	num := 0
	preSign := '+'
	for i, ch := range s {
		isD := isDigit(ch)
		if isD {
			num = num*10 + int(ch-'0')
		}
		if !isD && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			case '/':
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	res := 0
	for _, v := range stack {
		res += v
	}
	return res
}

func isDigit(s rune) bool {
	return s >= '0' && s <= '9'
}

// @lc code=end

