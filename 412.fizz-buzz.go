/*
 * @lc app=leetcode id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 0; i < n; i++ {
		res[i] = getString(i + 1)
	}
	return res
}

func getString(n int) string {
	if n%3 != 0 && n%5 != 0 {
		return strconv.Itoa(n)
	} else if n%3 == 0 && n%5 != 0 {
		return "Fizz"
	} else if n%3 != 0 && n%5 == 0 {
		return "Buzz"
	} else {
		return "FizzBuzz"
	}
}

// @lc code=end

