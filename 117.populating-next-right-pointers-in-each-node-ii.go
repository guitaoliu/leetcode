/*
 * @lc app=leetcode id=117 lang=golang
 *
 * [117] Populating Next Right Pointers in Each Node II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack
		stack = nil
		for i := 0; i < len(cur); i++ {
			if i < len(cur)-1 {
				cur[i].Next = cur[i+1]
			}
			if cur[i].Left != nil {
				stack = append(stack, cur[i].Left)
			}
			if cur[i].Right != nil {
				stack = append(stack, cur[i].Right)
			}
		}
	}
	return root
}

// @lc code=end

