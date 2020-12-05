/*
 * @lc app=leetcode id=507 lang=golang
 *
 * [507] Perfect Number
 */

// @lc code=start
import "math"

func checkPerfectNumber(num int) bool {
	if num == 1 {
		return false
	}
	sum := 1
	bound := int(math.Sqrt(float64(num)))
	for i := 2; i <= bound; i++ {
		if num%i == 0 {
			sum += num/i + i
		}
	}
	return num == sum
}

// @lc code=end

