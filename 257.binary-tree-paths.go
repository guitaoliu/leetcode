/*
 * @lc app=leetcode id=257 lang=golang
 *
 * [257] Binary Tree Paths
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
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	res := make([]string, 0)
	if root.Left == nil && root.Right == nil {
		return []string{strconv.Itoa(root.Val)}
	}
	tmpLeft := binaryTreePaths(root.Left)
	for _, v := range tmpLeft {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	tmpRight := binaryTreePaths(root.Right)
	for _, v := range tmpRight {
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	return res

}

// @lc code=end

