/*
 * @lc app=leetcode id=606 lang=golang
 *
 * [606] Construct String from Binary Tree
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
func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	left, right := "", ""
	if t.Left == nil && t.Right != nil {
		left = "()"
		right = "(" + tree2str(t.Right) + ")"
	} else if t.Left != nil && t.Right == nil {
		left = "(" + tree2str(t.Left) + ")"
	} else if t.Left != nil && t.Right != nil {
		left = "(" + tree2str(t.Left) + ")"
		right = "(" + tree2str(t.Right) + ")"
	}
	return strconv.Itoa(t.Val) + left + right
}

// @lc code=end

