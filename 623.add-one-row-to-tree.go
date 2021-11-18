/*
 * @lc app=leetcode id=623 lang=golang
 *
 * [623] Add One Row to Tree
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
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		node := &TreeNode{Val: val, Left: root}
		return node
	}
	return dfs(root, val, depth)
}

func dfs(root *TreeNode, val int, depth int) *TreeNode {
	if root == nil {
		return nil
	}
	if depth == 2 {
		if root.Left != nil {
			root.Left = &TreeNode{Val: val, Left: root.Left}
		} else {
			root.Left = &TreeNode{Val: val}
		}
		if root.Right != nil {
			root.Right = &TreeNode{Val: val, Right: root.Right}
		} else {
			root.Right = &TreeNode{Val: val}
		}
		return root
	}
	root.Left = dfs(root.Left, val, depth-1)
	root.Right = dfs(root.Right, val, depth-1)
	return root
}

// @lc code=end

