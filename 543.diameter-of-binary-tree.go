/*
 * @lc app=leetcode id=543 lang=golang
 *
 * [543] Diameter of Binary Tree
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

var res int = 0

func diameterOfBinaryTree(root *TreeNode) int {
	res = 0
	depth(root)
	return res
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	L, R := depth(root.Left), depth(root.Right)
	res = max(res, L+R)
	return max(L, R) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

