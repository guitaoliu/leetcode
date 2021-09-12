/*
 * @lc app=leetcode id=423 lang=golang
 *
 * [423] Reconstruct Original Digits from English
 */

// @lc code=start

func originalDigits(s string) string {
	digits := make([]int, 26)
	for i := 0; i < len(s); i++ {
		digits[s[i]-'a']++
	}
	res := make([]string, 10)
	res[0] = convert('z', digits, "zero", "0")
	res[2] = convert('w', digits, "two", "2")
	res[4] = convert('u', digits, "four", "4")
	res[6] = convert('x', digits, "six", "6")
	res[8] = convert('g', digits, "eight", "8")

	res[5] = convert('f', digits, "five", "5")
	res[7] = convert('s', digits, "seven", "7")
	res[3] = convert('h', digits, "three", "3")

	res[1] = convert('o', digits, "one", "1")
	res[9] = convert('i', digits, "nine", "9")
	return strings.Join(res, "")
}

func convert(k byte, digits []int, s string, num string) string {
	v := digits[k-'a']
	for i := 0; i < len(s); i++ {
		digits[s[i]-'a'] -= v
	}
	return strings.Repeat(num, v)
}

// @lc code=end

