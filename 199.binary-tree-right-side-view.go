/*
 * @lc app=leetcode id=199 lang=golang
 *
 * [199] Binary Tree Right Side View
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	res := []int{}
	for len(q) > 0 {
		res = append(res, q[len(q)-1].Val)
		curLevel := q
		q = nil
		for _, node := range curLevel {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}

// @lc code=end

