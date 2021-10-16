/*
 * @lc app=leetcode id=491 lang=golang
 *
 * [491] Increasing Subsequences
 */

// @lc code=start
func findSubsequences(nums []int) [][]int {
	res := [][]int{}
	var dfs func(int, []int)
	dfs = func(idx int, curList []int) {
		if len(curList) > 1 {
			tmp := make([]int, len(curList))
			copy(tmp, curList)
			res = append(res, tmp)
		}
		visited := make(map[int]struct{})
		for i := idx; i < len(nums); i++ {
			if len(curList) == 0 || nums[i] >= curList[len(curList)-1] {
				if _, ok := visited[nums[i]]; !ok {
					visited[nums[i]] = struct{}{}
					curList = append(curList, nums[i])
					dfs(i+1, curList)
					curList = curList[:len(curList)-1]
				}
			}
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		dfs(i, []int{})
	}
	return res
}

// @lc code=end

