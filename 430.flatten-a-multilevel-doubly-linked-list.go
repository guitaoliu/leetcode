/*
 * @lc app=leetcode id=430 lang=golang
 *
 * [430] Flatten a Multilevel Doubly Linked List
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}
	dummyHead := &Node{Next: root}
	prev := dummyHead

	stack := []*Node{root}
	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		prev.Next = cur
		cur.Prev = prev

		if cur.Next != nil {
			stack = append(stack, cur.Next)
		}

		if cur.Child != nil {
			stack = append(stack, cur.Child)
			cur.Child = nil
		}

		prev = cur
	}

	dummyHead.Next.Prev = nil
	return dummyHead.Next
}

// @lc code=end

