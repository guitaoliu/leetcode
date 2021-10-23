/*
 * @lc app=leetcode id=513 lang=golang
 *
 * [513] Find Bottom Left Tree Value
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
func findBottomLeftValueBFS(root *TreeNode) int {
	queue := []*TreeNode{root}
	res := 0
	for len(queue) > 0 {
		res = queue[0].Val
		curLevel := queue
		queue = nil
		for _, node := range curLevel {
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

type Node struct {
	treeNode *TreeNode
	level    int
}

func findBottomLeftValue(root *TreeNode) int {
	res := root.Val
	minLevel := 0
	stack := []*Node{&Node{treeNode: root, level: 0}}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.level > minLevel {
			res = node.treeNode.Val
			minLevel = node.level
		}
		if node.treeNode.Right != nil {
			stack = append(stack, &Node{node.treeNode.Right, node.level + 1})
		}
		if node.treeNode.Left != nil {
			stack = append(stack, &Node{node.treeNode.Left, node.level + 1})
		}
	}
	return res
}

// @lc code=end

