/*
 * @lc app=leetcode id=559 lang=golang
 *
 * [559] Maximum Depth of N-ary Tree
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	} else if root.Children == nil {
		return 1
	} else {
		return max(root.Children) + 1
	}
}

func max(nodes []*Node) int {
	max := 0
	for _, node := range nodes {
		depth := maxDepth(node)
		if depth > max {
			max = depth
		}
	}
	return max
}

// @lc code=end

