/*
 * @lc app=leetcode id=107 lang=golang
 *
 * [107] Binary Tree Level Order Traversal II
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
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	curNum, nextLevelNum, res, tmp := 1, 0, [][]int{}, []int{}
	for len(queue) > 0 {
		if curNum > 0 {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
				nextLevelNum++
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
				nextLevelNum++
			}
			curNum--
			tmp = append(tmp, node.Val)
			queue = queue[1:]
		}
		if curNum == 0 {
			res = append(res, tmp)
			curNum, nextLevelNum = nextLevelNum, 0
			tmp = []int{}
		}
	}
	return reverse(res)
}

func reverseIntSlices(input [][]int) [][]int {
	if len(input) == 0 {
		return input
	}
	return append(reverseIntSlices(input[1:]), input[0])
}

func reverse(input [][]int) [][]int {
	for i := 0; i < len(input)/2; i++ {
		j := len(input) - i - 1
		input[i], input[j] = input[j], input[i]
	}
	return input
}

// @lc code=end

