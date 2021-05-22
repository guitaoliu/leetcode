/*
 * @lc app=leetcode id=328 lang=golang
 *
 * [328] Odd Even Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	odd, even := head, head.Next
	oddPtr, evenPtr := odd, even
	head = head.Next.Next
	isOdd := true
	for head != nil {
		if isOdd {
			oddPtr.Next = head
			oddPtr = oddPtr.Next
		} else {
			evenPtr.Next = head
			evenPtr = evenPtr.Next
		}
		isOdd = !isOdd
		head = head.Next
	}
	oddPtr.Next = even
	evenPtr.Next = nil
	return odd
}

// @lc code=end

