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
	start := root
	for start != nil {
		var next, last *Node
		handle := func(node *Node) {
			if node == nil {
				return
			}
			if next == nil {
				next = node
			}
			if last != nil {
				last.Next = node
			}
			last = node
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = next
	}
	return root
}

// @lc code=end

