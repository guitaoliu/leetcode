/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if isSameTree(s, t) {
		return true
	}
	if s == nil {
		return false
	}
	if isSubtree(s.Left, t) || isSubtree(s.Right, t) {
		return true
	}
	return false
}

func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	} else if s != nil && t != nil {
		if s.Val != t.Val {
			return false
		}
		return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
	} else {
		return false
	}
}

// @lc code=end

