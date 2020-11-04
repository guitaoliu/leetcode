/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
// func isAnagram(s string, t string) bool {
// 	oriMap := make(map[byte]int, 0)
// 	sBytes := []byte(s)
// 	tBytes := []byte(t)
// 	for _, v := range sBytes {
// 		if _, ok := oriMap[v]; ok {
// 			oriMap[v]++
// 		} else {
// 			oriMap[v] = 1
// 		}
// 	}
// 	for _, v := range tBytes {
// 		if _, ok := oriMap[v]; ok {
// 			if oriMap[v] > 0 {
// 				oriMap[v]--
// 			} else {
// 				return false
// 			}
// 		} else {
// 			return false
// 		}
// 	}
// 	for _, v := range oriMap {
// 		if v != 0 {
// 			return false
// 		}
// 	}
// 	return true
// }

func isAnagram(s string, t string) bool {
	alphabet := make([]int, 26)
	sBytes := []byte(s)
	tBytes := []byte(t)
	if len(sBytes) != len(tBytes) {
		return false
	}
	for i := 0; i < len(sBytes); i++ {
		alphabet[sBytes[i]-'a']++
	}
	for i := 0; i < len(tBytes); i++ {
		alphabet[tBytes[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if alphabet[i] != 0 {
			return false
		}
	}
	return true
}

// @lc code=end

