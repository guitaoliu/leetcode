/*
 * @lc app=leetcode id=421 lang=golang
 *
 * [421] Maximum XOR of Two Numbers in an Array
 */

// @lc code=start
func findMaximumXORBruteForce(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			tmp := nums[i] ^ nums[j]
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

func findMaximumXOR(nums []int) int {
	max, mask := 0, 0
	for i := 31; i >= 0; i-- {
		// 按照 100...000, 110...000 增长
		mask = mask | (1 << uint(i))
		m := make(map[int]struct{})
		for _, num := range nums {
			m[num&mask] = struct{}{}
		}

		greedyTry := max | (1 << uint(i))
		for anotherNum := range m {
			if _, ok := m[anotherNum^greedyTry]; ok {
				max = greedyTry
				break
			}
		}
	}

	return max
}

// @lc code=end

