/*
 * @lc app=leetcode id=671 lang=golang
 *
 * [671] Second Minimum Node In a Binary Tree
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
func findSecondMinimumValue(root *TreeNode) int {
	min, res := root.Val, math.MaxInt64
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		if node.Val < min {
			min, res = node.Val, min
		}
		if node.Val > min && node.Val < res {
			res = node.Val
		}
	}
	dfs(root)
	if res == math.MaxInt64 {
		return -1
	}
	return res
}

// @lc code=end

