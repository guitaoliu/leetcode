/*
 * @lc app=leetcode id=241 lang=golang
 *
 * [241] Different Ways to Add Parentheses
 */

// @lc code=start
func diffWaysToCompute(expression string) []int {
	if num, err := strconv.Atoi(expression); err == nil {
		return []int{num}
	}
	res := make([]int, 0)
	for idx, ch := range []byte(expression) {
		if ch == '*' || ch == '-' || ch == '+' {
			left := diffWaysToCompute(expression[:idx])
			right := diffWaysToCompute(expression[idx+1:])
			operator := getOperation(ch)
			for _, l := range left {
				for _, r := range right {
					num := operator(l, r)
					res = append(res, num)
				}
			}
		}
	}
	return res
}

func getOperation(ch byte) func(int, int) int {
	if ch == '*' {
		return func(a, b int) int { return a * b }
	} else if ch == '+' {
		return func(a, b int) int { return a + b }
	} else {
		return func(a, b int) int { return a - b }
	}
}

// @lc code=end

