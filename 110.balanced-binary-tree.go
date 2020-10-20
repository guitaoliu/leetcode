/*
 * @lc app=leetcode id=110 lang=golang
 *
 * [110] Balanced Binary Tree
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
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return leftDepth-rightDepth <= 1 &&
		leftDepth-rightDepth >= -1 &&
		isBalanced(root.Left) &&
		isBalanced(root.Right)

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1

}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// @lc code=end

