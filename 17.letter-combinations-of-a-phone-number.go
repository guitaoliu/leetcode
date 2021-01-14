/*
 * @lc app=leetcode id=17 lang=golang
 *
 * [17] Letter Combinations of a Phone Number
 */

// @lc code=start

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	m := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	res := make([]string, 0)
	digitsByte := []byte(digits)
	var backtrace func(idx int, letter []byte)
	backtrace = func(idx int, letter []byte) {
		if idx == len(digitsByte) {
			res = append(res, string(letter))
		} else {
			digit := digitsByte[idx]
			for _, v := range m[digit] {
				backtrace(idx+1, append(letter, v))
			}
		}
	}
	backtrace(0, []byte{})
	return res
}

// @lc code=end

