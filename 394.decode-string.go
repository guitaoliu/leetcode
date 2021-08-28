/*
 * @lc app=leetcode id=394 lang=golang
 *
 * [394] Decode String
 */

// @lc code=start
func decodeString(s string) string {
	stack := []byte{}
	sBytes := []byte(s)
	for _, b := range sBytes {
		if b == ']' {
			tmp := []byte{}
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				tmp = append(tmp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			tmp = reverse(tmp)

			num := []byte{}
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				num = append(num, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			times, _ := strconv.Atoi(string(reverse(num)))

			timesTmp := []byte{}
			for j := 0; j < times; j++ {
				timesTmp = append(timesTmp, tmp...)
			}

			stack = append(stack, timesTmp...)
		} else {
			stack = append(stack, b)
		}
	}

	return string(stack)
}

func reverse(list []byte) []byte {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
	return list
}

// @lc code=end

