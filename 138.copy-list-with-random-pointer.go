/*
 * @lc app=leetcode id=138 lang=golang
 *
 * [138] Copy List with Random Pointer
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodeMap := make(map[*Node]*Node, 0)
	var clone func(*Node) *Node
	clone = func(node *Node) *Node {
		if node == nil {
			return nil
		}
		if n, ok := nodeMap[node]; ok {
			return n
		}
		newNode := &Node{Val: node.Val}
		nodeMap[node] = newNode
		newNode.Next = clone(node.Next)
		newNode.Random = clone(node.Random)
		return newNode
	}
	return clone(head)
}

// @lc code=end

