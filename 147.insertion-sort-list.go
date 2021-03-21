/*
 * @lc app=leetcode id=147 lang=golang
 *
 * [147] Insertion Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -5001, Next: head}
	cur := dummyHead
	insert := func(node *ListNode) {
		p := dummyHead
		for p.Next.Val < node.Val {
			p = p.Next
		}
		node.Next = p.Next
		p.Next = node
	}
	for cur.Next != nil {
		for cur.Next != nil && cur.Next.Val > cur.Val {
			cur = cur.Next
		}
		if cur.Next != nil {
			tmp := cur.Next
			cur.Next = cur.Next.Next
			insert(tmp)
		}
	}
	return dummyHead.Next
}

// @lc code=end

