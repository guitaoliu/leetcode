/*
 * @lc app=leetcode id=258 lang=golang
 *
 * [258] Add Digits
 */

// @lc code=start
// func addDigits(num int) int {
// 	for num > 9 {
// 		cur := 0
// 		for num != 0 {
// 			cur += num % 10
// 			num /= 10
// 		}
// 		num = cur
// 	}
// 	return num
// }

// func addDigits(num int) int {
// 	if num == 0 {
// 		return 0
// 	} else if num%9 == 0 {
// 		return 9
// 	} else {
// 		return num % 9
// 	}
// }

func addDigits(num int) int {
	return (num-1)%9 + 1
}

// @lc code=end

