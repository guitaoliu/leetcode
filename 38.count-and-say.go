/*
 * @lc app=leetcode id=38 lang=golang
 *
 * [38] Count and Say
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	strNext := countAndSay(n - 1)
	c := strNext[0]
	cnt := 1
	strCurrent := ""

	for i := 1; i < len(strNext); i++ {
		if strNext[i] == c {
			cnt++
			continue
		}
		strCurrent += strconv.Itoa(cnt) + string(c)
		c = strNext[i]
		cnt = 1
	}
	strCurrent += strconv.Itoa(cnt) + string(c)
	return strCurrent
}

// @lc code=end


