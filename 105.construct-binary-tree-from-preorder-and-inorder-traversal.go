/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	inorderPos := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderPos[v] = i
	}
	return build(preorder, 0, len(preorder)-1, 0, inorderPos)
}

func build(preorder []int, start, end, inorderStart int, inorderPos map[int]int) *TreeNode {
	if start > end {
		return nil
	}
	root := &TreeNode{Val: preorder[start]}
	if pos, ok := inorderPos[root.Val]; ok {
		leftLen := pos - inorderStart
		root.Left = build(preorder, start+1, start+leftLen, inorderStart, inorderPos)
		root.Right = build(preorder, start+leftLen+1, end, pos+1, inorderPos)
	}
	return root
}

// @lc code=end

