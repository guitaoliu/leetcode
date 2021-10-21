/*
 * @lc app=leetcode id=508 lang=golang
 *
 * [508] Most Frequent Subtree Sum
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findFrequentTreeSum(root *TreeNode) []int {
	m := make(map[int]int)
	var trace func(*TreeNode) int
	trace = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		sum := trace(node.Left) + trace(node.Right) + node.Val
		m[sum]++
		return sum
	}
	trace(root)
	max := 0
	for _, freq := range m {
		if freq > max {
			max = freq
		}
	}
	res := []int{}
	for sum, freq := range m {
		if freq == max {
			res = append(res, sum)
		}
	}
	return res
}

// @lc code=end

