/*
 * @lc app=leetcode id=637 lang=golang
 *
 * [637] Average of Levels in Binary Tree
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

func averageOfLevels(root *TreeNode) []float64 {
	curLevel := []*TreeNode{root}
	res := []float64{}
	for len(curLevel) > 0 {
		sum := 0
		queue := make([]*TreeNode, 0)
		for _, node := range curLevel {
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, float64(sum)/float64(len(curLevel)))
		curLevel = queue
	}
	return res
}

// @lc code=end

