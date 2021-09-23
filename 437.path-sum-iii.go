/*
 * @lc app=leetcode id=437 lang=golang
 *
 * [437] Path Sum III
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
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	prefixSum := make(map[int]int)
	prefixSum[0] = 1
	res := pathSumPrefixSumHelper(root, prefixSum, 0, targetSum)
	return res
}

func pathSumPrefixSumHelper(node *TreeNode, prefixSum map[int]int, cur int, targetSum int) int {
	if node == nil {
		return 0
	}
	cur += node.Val
	cnt := 0
	if v, ok := prefixSum[cur-targetSum]; ok {
		cnt = v
	}
	prefixSum[cur]++
	cnt += pathSumPrefixSumHelper(node.Left, prefixSum, cur, targetSum)
	cnt += pathSumPrefixSumHelper(node.Right, prefixSum, cur, targetSum)
	prefixSum[cur]--
	return cnt
}

func pathSumDFS(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	cnt := 0
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		for _, node := range curLevel {
			if paths := pathSumHelper(node, targetSum); paths > 0 {
				cnt += paths
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return cnt
}

func pathSumHelper(node *TreeNode, remain int) int {
	res := 0
	if node.Val == remain {
		res++
	}
	if node.Left != nil {
		res += pathSumHelper(node.Left, remain-node.Val)
	}
	if node.Right != nil {
		res += pathSumHelper(node.Right, remain-node.Val)
	}
	return res
}

// @lc code=end

