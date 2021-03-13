/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

var newNodes map[int]*Node

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNodes = make(map[int]*Node, 0)
	res := clone(node)
	return res
}

func clone(node *Node) *Node {
	if n, ok := newNodes[node.Val]; ok {
		return n
	}
	newNode := &Node{Val: node.Val}
	newNodes[newNode.Val] = newNode
	for _, nei := range node.Neighbors {
		newNode.Neighbors = append(newNode.Neighbors, clone(nei))
	}
	return newNode
}

// @lc code=end

