/*
 * @lc app=leetcode id=653 lang=golang
 *
 * [653] Two Sum IV - Input is a BST
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
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool, 0)
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		curLevel := stack
		stack = nil
		for _, v := range curLevel {
			if _, ok := m[k-v.Val]; ok {
				return true
			} else {
				m[v.Val] = true
			}
			if v.Left != nil {
				stack = append(stack, v.Left)
			}
			if v.Right != nil {
				stack = append(stack, v.Right)
			}
		}
	}
	return false
}

// @lc code=end

