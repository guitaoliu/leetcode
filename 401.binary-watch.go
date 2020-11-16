/*
 * @lc app=leetcode id=401 lang=golang
 *
 * [401] Binary Watch
 */

// @lc code=start
import "strconv"

func readBinaryWatch(num int) []string {
	res := []string{}
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if count(i)+count(j) == num {
				if j < 10 {
					res = append(res, strconv.Itoa(i)+":0"+strconv.Itoa(j))
				} else {
					res = append(res, strconv.Itoa(i)+":"+strconv.Itoa(j))
				}
			}
		}
	}
	return res
}

func count(n int) int {
	res := 0
	for n > 0 {
		n &= n - 1
		res++
	}
	return res
}

// @lc code=end

