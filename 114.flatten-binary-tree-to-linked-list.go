/*
 * @lc app=leetcode id=114 lang=golang
 *
 * [114] Flatten Binary Tree to Linked List
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

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tmp := root.Right
	root.Right = root.Left
	root.Left = nil
	for root.Right != nil {
		root = root.Right
	}
	root.Right = tmp
}

// func flatten(root *TreeNode) {
// 	cur := root
// 	for cur != nil {
// 		if cur.Left != nil {
// 			next := cur.Left
// 			pre := next
// 			for pre.Right != nil {
// 				pre = pre.Right
// 			}
// 			pre.Right = cur.Right
// 			cur.Left, cur.Right = nil, next
// 		}
// 		cur = cur.Right
// 	}
// }

// @lc code=end

