/*
 * @lc app=leetcode id=86 lang=golang
 *
 * [86] Partition List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	greatHead := &ListNode{Val: 0, Next: nil}
	gCur := greatHead
	lowHead := &ListNode{Val: 0, Next: nil}
	lCur := lowHead
	cur := head
	for cur != nil {
		if cur.Val >= x {
			gCur.Next = cur
			gCur = gCur.Next
		} else {
			lCur.Next = cur
			lCur = lCur.Next
		}
		cur = cur.Next
	}
	gCur.Next = nil
	lCur.Next = greatHead.Next
	return lowHead.Next
}

// @lc code=end

