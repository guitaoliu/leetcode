/*
 * @lc app=leetcode id=515 lang=golang
 *
 * [515] Find Largest Value in Each Tree Row
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
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	levelOrder := [][]int{}
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		tmp := []int{}
		for _, node := range curLevel {
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, tmp)
	}
	res := make([]int, len(levelOrder))
	for i, level := range levelOrder {
		sort.Ints(level)
		res[i] = level[len(level)-1]
	}
	return res
}

func largestValuesBFS(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		max := math.MinInt64
		for _, node := range curLevel {
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, max)
	}
	return res
}

// @lc code=end

