/*
 * @lc app=leetcode id=563 lang=golang
 *
 * [563] Binary Tree Tilt
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
func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	dfs(root, &res)
	return res
}

func dfs(node *TreeNode, sum *int) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left, sum)
	right := dfs(node.Right, sum)
	*sum += abs(left - right)
	return node.Val + left + right
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

// @lc code=end

