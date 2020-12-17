/*
 * @lc app=leetcode id=590 lang=golang
 *
 * [590] N-ary Tree Postorder Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	res := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		stack = append(stack, node.Children...)
	}
	return res
}

// @lc code=end

