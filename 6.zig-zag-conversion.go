/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	rows := make([][]byte, numRows)
	for i := 0; i < len(rows); i++ {
		rows[i] = make([]byte, 0)
	}
	curRow := 0
	down := true
	for i := 0; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])
		if down {
			curRow++
		} else {
			curRow--
		}

		if curRow == 0 {
			down = true
		}
		if curRow == numRows-1 {
			down = false
		}
	}
	res := make([]byte, 0, len(s))
	for i := 0; i < numRows; i++ {
		res = append(res, rows[i]...)
	}
	return string(res)
}

// @lc code=end

