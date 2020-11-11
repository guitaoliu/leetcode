/*
 * @lc app=leetcode id=342 lang=golang
 *
 * [342] Power of Four
 */

// @lc code=start
import "math"

func isPowerOfFour(num int) bool {
	powerMap := powers()
	_, ok := powerMap[num]
	return ok
}

func powers() map[int]int {
	res := make(map[int]int, 15)
	for i := 1; i <= 15; i++ {
		res[i] = int(math.Pow(4, float64(i)))
	}
	return res
}

// @lc code=end

