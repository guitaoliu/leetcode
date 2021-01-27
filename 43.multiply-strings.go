/*
 * @lc app=leetcode id=43 lang=golang
 *
 * [43] Multiply Strings
 */

// @lc code=start
func multiply(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	num1Bytes := []byte(num1)
	num2Bytes := []byte(num2)
	res := make([]byte, 0)
	idxNum1 = len(num1Bytes) - 1
	idxNum2 = len(num2Bytes) - 1
	c := 0
	for idxNum2 >= 0 {
		tmp = int(num1Bytes[idxNum1]-'0')*int(num2Bytes[idxNum2]-'0') + c
		if tmp < 10 {
			res = append(res, byte(tmp+'0'))
		} else {
			c = tmp % 10
			tmp /= 10
			res = append(res, byte(tmp+'0'))
		}
		idxNum1 --
		idxNum2 --
	}
	for 
	return res
}

// @lc code=end

