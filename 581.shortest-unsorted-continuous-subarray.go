/*
 * @lc app=leetcode id=581 lang=golang
 *
 * [581] Shortest Unsorted Continuous Subarray
 */

// @lc code=start
func findUnsortedSubarrayTwoPointers(nums []int) int {
	maxNum := math.MinInt64
	minNum := math.MaxInt64
	n := len(nums)
	l, r := n-1, 0
	for i := 0; i < n; i++ {
		if nums[i] < maxNum {
			r = i
		}
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
		if nums[n-i-1] > minNum {
			l = n - i - 1
		}
		if nums[n-i-1] < minNum {
			minNum = nums[n-i-1]
		}
	}
	if l == n-1 && r == 0 {
		return 0
	}
	return r - l + 1
}

func findUnsortedSubarray(nums []int) int {
	tmp := make([]int, len(nums))
	copy(tmp, nums)
	sort.Ints(tmp)
	i, j := 0, len(nums)-1
	for i <= j && nums[i] == tmp[i] {
		i++
	}
	for i <= j && nums[j] == tmp[j] {
		j--
	}
	return j - i + 1
}

// @lc code=end

