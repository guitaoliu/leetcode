/*
 * @lc app=leetcode id=173 lang=golang
 *
 * [173] Binary Search Tree Iterator
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
type BSTIterator struct {
	arr []int
	cur int
}

func Constructor(root *TreeNode) BSTIterator {
	it := BSTIterator{arr: make([]int, 0)}
	it.inorder(root)
	return it
}

func (this *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}
	this.inorder(node.Left)
	this.arr = append(this.arr, node.Val)
	this.inorder(node.Right)
}

func (this *BSTIterator) Next() int {
	ret := this.arr[this.cur]
	this.cur++
	return ret
}

func (this *BSTIterator) HasNext() bool {
	return this.cur < len(this.arr)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
// @lc code=end

