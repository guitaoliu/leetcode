/*
 * @lc app=leetcode id=109 lang=golang
 *
 * [109] Convert Sorted List to Binary Search Tree
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	return build(head, nil)
}

func build(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getLinkedListMid(left, right)
	return &TreeNode{
		Val:   mid.Val,
		Left:  build(left, mid),
		Right: build(mid.Next, right),
	}
}

func getLinkedListMid(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// @lc code=end

