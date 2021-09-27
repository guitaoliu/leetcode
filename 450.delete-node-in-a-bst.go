/*
 * @lc app=leetcode id=450 lang=golang
 *
 * [450] Delete Node in a BST
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
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil && root.Right == nil {
			root = nil
		} else if root.Right != nil {
			root.Val = rightMin(root)
			root.Right = deleteNode(root.Right, root.Val)
		} else {
			root.Val = leftMax(root)
			root.Left = deleteNode(root.Left, root.Val)
		}
	}
	return root
}

func rightMin(node *TreeNode) int {
	node = node.Right
	for node.Left != nil {
		node = node.Left
	}
	return node.Val
}

func leftMax(node *TreeNode) int {
	node = node.Left
	for node.Right != nil {
		node = node.Right
	}
	return node.Val
}

// @lc code=end

