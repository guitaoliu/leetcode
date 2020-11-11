/*
 * @lc app=leetcode id=345 lang=golang
 *
 * [345] Reverse Vowels of a String
 */

// @lc code=start
func reverseVowels(s string) string {
	vowelsMap := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
		'A': 1,
		'E': 1,
		'I': 1,
		'O': 1,
		'U': 1,
	}
	i, j := 0, len(s)-1
	sByte := []byte(s)
	for i < j {
		if _, ok := vowelsMap[sByte[i]]; !ok {
			i++
		} else if _, ok := vowelsMap[sByte[j]]; !ok {
			j--
		} else {
			sByte[i], sByte[j] = sByte[j], sByte[i]
			i++
			j--
		}
	}
	return string(sByte)
}

// @lc code=end

