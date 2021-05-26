/*
 * @lc app=leetcode id=337 lang=golang
 *
 * [337] House Robber III
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
func rob(root *TreeNode) int {
	a, b := dfs(root)
	return max(a, b)
}

func dfs(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}
	la, lb := dfs(root.Left)
	ra, rb := dfs(root.Right)
	a = max(la, lb) + max(ra, rb)
	b = root.Val + la + ra
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

