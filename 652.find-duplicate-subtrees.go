/*
 * @lc app=leetcode id=652 lang=golang
 *
 * [652] Find Duplicate Subtrees
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

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	ids := make(map[int]int)
	counts := make(map[int]int)

	res := []*TreeNode{}

	var marshal func(*TreeNode) int
	marshal = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		uid := node.Val<<32 | marshal(node.Left)<<16 | marshal(node.Right)
		if _, ok := ids[uid]; !ok {
			ids[uid] = len(ids) + 1
		}
		counts[ids[uid]]++
		if counts[ids[uid]] == 2 {
			res = append(res, node)
		}
		return ids[uid]
	}
	marshal(root)
	return res
}

func findDuplicateSubtreesMarshal(root *TreeNode) []*TreeNode {
	m := make(map[string]int)
	res := []*TreeNode{}

	var marshal func(*TreeNode) string
	marshal = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		tmp := fmt.Sprintf("%v,%v,%v", node.Val, marshal(node.Left), marshal(node.Right))
		m[tmp]++
		if m[tmp] == 2 {
			res = append(res, node)
		}
		return tmp
	}
	marshal(root)
	return res
}

// @lc code=end

