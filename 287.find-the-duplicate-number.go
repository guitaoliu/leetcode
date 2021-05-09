/*
 * @lc app=leetcode id=287 lang=golang
 *
 * [287] Find the Duplicate Number
 */

// @lc code=start

func findDuplicate(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		cnt := 0
		for _, num := range nums {
			if num <= mid {
				cnt++
			}
		}
		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func findDuplicateTwoPointer(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for fast != slow {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	res := 0
	for res != slow {
		res = nums[res]
		slow = nums[slow]
	}
	return res
}

func findDuplicateHash(nums []int) int {
	n := len(nums) - 1
	m := make([]int, n)
	for _, num := range nums {
		if m[num-1] > 0 {
			return num
		}
		m[num-1]++
	}
	return n
}

// @lc code=end

