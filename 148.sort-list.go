/*
 * @lc app=leetcode id=148 lang=golang
 *
 * [148] Sort List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow.Next
	slow.Next = nil
	left := sortList(head)
	right := sortList(mid)
	return mergeList(left, right)
}

func mergeList(h1, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	if h1.Val > h2.Val {
		h2.Next = mergeList(h1, h2.Next)
		return h2
	}
	h1.Next = mergeList(h1.Next, h2)
	return h1
}

// @lc code=end

