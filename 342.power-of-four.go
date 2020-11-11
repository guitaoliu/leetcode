/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
func isPowerOfFour(num int) bool {
	powerMap := powers()
	_, ok := powerMap[num]
	return ok
}

func powers() map[int]int {
	res := make(map[int]int, 15)
	for i := 0; i <= 15; i++ {
		res[powerOfFour(i)] = i
	}
	return res
}

func powerOfFour(n int) int {
	res := 1
	for n > 0 {
		res *= 4
		n--
	}
	return res
}

// @lc code=end

