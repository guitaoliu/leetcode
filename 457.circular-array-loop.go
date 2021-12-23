/*
 * @lc app=leetcode id=457 lang=golang
 *
 * [457] Circular Array Loop
 */

// @lc code=start
func circularArrayLoop(nums []int) bool {
	for i := range nums {
		direction := nums[i] > 0
		fast, slow := i, i
		for {
			slow = next(nums, direction, slow)
			fast = next(nums, direction, fast)
			if fast != -1 {
				fast = next(nums, direction, fast)
			}
			if fast == -1 || slow == -1 || fast == slow {
				break
			}
		}
		if slow != -1 && fast == slow {
			return true
		}
	}
	return false
}

func next(nums []int, isForward bool, index int) int {
	direction := nums[index] > 0
	if isForward != direction {
		return -1
	}

	nextIndex := ((index+nums[index])%len(nums) + len(nums)) % len(nums)
	if nextIndex == index {
		return -1
	}
	return nextIndex
}

// @lc code=end

