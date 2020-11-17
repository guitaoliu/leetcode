/*
 * @lc app=leetcode id=405 lang=golang
 *
 * [405] Convert a Number to Hexadecimal
 */

// @lc code=start

func toHex(num int) string {
	if num == 0 {
		return "0"
	}
	if num < 0 {
		num += 1 << 32
	}
	m := map[int]byte{
		0: '0', 1: '1', 2: '2', 3: '3', 4: '4', 5: '5', 6: '6', 7: '7', 8: '8', 9: '9',
		10: 'a', 11: 'b', 12: 'c', 13: 'd', 14: 'e', 15: 'f',
	}
	res := []byte{}
	for num > 0 {
		res = append(res, m[num%16])
		num /= 16
	}
	for num < 0 {

	}
	return string(reverse(res))
}

func reverse(res []byte) []byte {
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// @lc code=end

