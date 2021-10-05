/*
 * @lc app=leetcode id=462 lang=golang
 *
 * [462] Minimum Moves to Equal Array Elements II
 */

// @lc code=start

func minMoves2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	res := 0
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		res += nums[j] - nums[i]
	}
	return res
}

func minMoves2Mid(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	mid := nums[len(nums)/2]
	res := 0
	for _, num := range nums {
		res += abs(mid - num)
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end

