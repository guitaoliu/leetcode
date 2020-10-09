/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */

// @lc code=start
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a, b = b, a
	}

	res := make([]string, len(a))
	j, k := len(a)-1, len(b)-1
	c := 0
	result := 0
	for j >= 0 {
		aInt, _ := strconv.Atoi(string(a[j]))
		if k >= 0 {
			bInt, _ := strconv.Atoi(string(b[k]))
			result = (bInt + c + aInt) % 2
			c = (bInt + c + aInt) / 2
		} else {
			result = (c + aInt) % 2
			c = (c + aInt) / 2
		}

		res[j] = strconv.Itoa(result)
		j--
		k--
	}
	if c > 0 {
		res = append([]string{"1"}, res...)
	}
	return strings.Join(res, "")
}

// @lc code=end

