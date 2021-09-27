/*
 * @lc app=leetcode id=449 lang=golang
 *
 * [449] Serialize and Deserialize BST
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
import (
	"strconv"
)

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "nil"
	}
	res := []string{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevel := queue
		queue = nil
		for _, node := range curLevel {
			if node != nil {
				res = append(res, strconv.Itoa(node.Val))
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			} else {
				res = append(res, "nil")
			}
		}
	}
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nodeVals := strings.Split(data, ",")
	val, err := strconv.Atoi(nodeVals[0])
	if err != nil {
		return nil
	}
	root := &TreeNode{Val: val}
	nodeVals = nodeVals[1:]
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curLevel := queue
		curLevelLen := len(curLevel)
		queue = nil
		for i := 0; i < curLevelLen; i++ {
			leftVal, err := strconv.Atoi(nodeVals[0])
			if err != nil {
				curLevel[i].Left = nil
			} else {
				node := &TreeNode{Val: leftVal}
				curLevel[i].Left = node
				queue = append(queue, node)
			}

			rightVal, err := strconv.Atoi(nodeVals[1])
			if err != nil {
				curLevel[i].Right = nil
			} else {
				node := &TreeNode{Val: rightVal}
				curLevel[i].Right = node
				queue = append(queue, node)
			}

			nodeVals = nodeVals[2:]
		}
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
// @lc code=end

