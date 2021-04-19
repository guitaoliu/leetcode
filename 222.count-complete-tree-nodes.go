/*
 * @lc app=leetcode id=222 lang=golang
 *
 * [222] Count Complete Tree Nodes
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

func countNodesDFS(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := root.Left, root.Right
	lH, rH := 0, 0
	for l != nil {
		lH++
		l = l.Left
	}
	for r != nil {
		rH++
		r = r.Right
	}
	if lH == rH {
		return 2<<lH - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

// @lc code=end

