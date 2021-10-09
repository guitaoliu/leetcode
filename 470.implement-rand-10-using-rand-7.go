/*
 * @lc app=leetcode id=470 lang=golang
 *
 * [470] Implement Rand10() Using Rand7()
 */

// @lc code=start
func rand10() int {
	var res int
	for {
		res = (rand7()-1)*7 + rand7()
		if res <= 40 {
			return res%10 + 1
		}

		res -= 40
		res = (res-1)*7 + rand7()
		if res <= 60 {
			return res%10 + 1
		}

		res -= 60
		res = (res-1)*7 + rand7()
		if res <= 20 {
			return res%10 + 1
		}

	}
}

// @lc code=end

