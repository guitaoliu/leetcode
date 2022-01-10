/*
 * @lc app=leetcode id=698 lang=golang
 *
 * [698] Partition to K Equal Sum Subsets
 */

// @lc code=start
func canPartitionKSubsets(nums []int, k int) bool {
	if k > len(nums) {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	sort.Ints(nums)
	return dfs(nums, sum/k, 0, k, 0)
}

func dfs(nums []int, target, cur, remain, used int) bool {
	if remain == 0 {
		return used == 1<<len(nums)-1
	}
	for i, num := range nums {
		if used&1<<i != 0 {
			continue
		}
		tmp := cur + num
		if tmp > target {
			break
		}
		if tmp == target && dfs(nums, target, 0, remain-1, used|1<<i) {
			return true
		} else if dfs(nums, target, tmp, remain, used|1<<i) {
			return true
		}
	}
	return false
}

// @lc code=end

