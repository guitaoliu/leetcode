/*
 * @lc app=leetcode id=113 lang=golang
 *
 * [113] Path Sum II
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

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var trace func(*TreeNode, int)
	trace = func(node *TreeNode, remaining int) {
		if node == nil {
			return
		}
		remaining -= node.Val
		path = append(path, node.Val)
		defer func() {
			path = path[:len(path)-1]
		}()
		if node.Left == nil && node.Right == nil && remaining == 0 {
			res = append(res, append([]int{}, path...))
			return
		}
		trace(node.Left, remaining)
		trace(node.Right, remaining)
	}
	trace(root, targetSum)
	return res
}

// @lc code=end

