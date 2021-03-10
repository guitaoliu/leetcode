/*
 * @lc app=leetcode id=129 lang=golang
 *
 * [129] Sum Root to Leaf Numbers
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
func sumNumbers(root *TreeNode) int {
	res := 0
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, num string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			sum, _ := strconv.Atoi(num + strconv.Itoa(node.Val))
			res += sum
		} else {
			dfs(node.Left, num+strconv.Itoa(node.Val))
			dfs(node.Right, num+strconv.Itoa(node.Val))
		}
	}
	dfs(root, "")
	return res
}

// @lc code=end

