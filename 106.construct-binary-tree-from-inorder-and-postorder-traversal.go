/*
 * @lc app=leetcode id=106 lang=golang
 *
 * [106] Construct Binary Tree from Inorder and Postorder Traversal
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
var inorderPos map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	inorderPos = make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderPos[v] = i
	}
	return build(postorder, 0, len(postorder)-1, 0)
}

func build(postorder []int, start, end int, inPos int) *TreeNode {
	if start > end {
		return nil
	}
	root := &TreeNode{Val: postorder[end]}
	if pos, ok := inorderPos[root.Val]; ok {
		lenLeft := pos - inPos
		root.Left = build(postorder, start, start+lenLeft-1, inPos)
		root.Right = build(postorder, start+lenLeft, end-1, pos+1)
	}
	return root
}

// @lc code=end

