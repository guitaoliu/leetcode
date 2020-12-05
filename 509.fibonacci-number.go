/*
 * @lc app=leetcode id=509 lang=golang
 *
 * [509] Fibonacci Number
 */

// @lc code=start
func fib(N int) int {
	res := 0
	tmp1, tmp2 := 0, 1
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	for i := 1; i < N; i++ {
		res = tmp1 + tmp2
		tmp1, tmp2 = tmp2, res
	}
	return res
}

// @lc code=end

