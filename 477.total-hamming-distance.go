/*
 * @lc app=leetcode id=477 lang=golang
 *
 * [477] Total Hamming Distance
 */

// @lc code=start

func totalHammingDistance(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		bitCnt := 0
		for j := 0; j < len(nums); j++ {
			bitCnt += (nums[j] >> i) & 1
		}
		res += bitCnt * (len(nums) - bitCnt)
	}
	return res
}

func totalHammingDistanceBrute(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res += count(nums[i] ^ nums[j])
		}
	}
	return res
}

func count(n int) int {
	cnt := 0
	for n > 0 {
		cnt += n & 1
		n = n >> 1
	}
	return cnt
}

// @lc code=end

