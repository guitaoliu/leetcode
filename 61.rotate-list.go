/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	length := 1
	cur := head
	for cur.Next != nil {
		length++
		cur = cur.Next
	}
	cur.Next = head
	cur = head
	for v := length - k%length - 1; v > 0; v-- {
		cur = cur.Next
	}
	res := cur.Next
	cur.Next = nil
	return res
}

// @lc code=end

