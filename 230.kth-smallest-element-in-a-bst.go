/*
 * @lc app=leetcode id=230 lang=golang
 *
 * [230] Kth Smallest Element in a BST
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
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	cnt := 0
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		cnt++
		if cnt == k {
			res = node.Val
			return
		}
		inorder(node.Right)
	}
	inorder(root)
	return res
}

// @lc code=end

