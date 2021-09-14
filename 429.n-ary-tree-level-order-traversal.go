/*
 * @lc app=leetcode id=429 lang=golang
 *
 * [429] N-ary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*Node{root}
	res := [][]int{}
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		tmp := []int{}
		for _, node := range curLevel {
			tmp = append(tmp, node.Val)
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		res = append(res, tmp)
	}
	return res
}

// @lc code=end

