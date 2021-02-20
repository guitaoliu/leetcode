/*
 * @lc app=leetcode id=82 lang=golang
 *
 * [82] Remove Duplicates from Sorted List II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			tmp := cur.Next
			for tmp.Next != nil && tmp.Next.Val == tmp.Val {
				tmp = tmp.Next
			}
			cur.Next = tmp.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// func deleteDuplicates(head *ListNode) *ListNode {
// 	if head == nil || head.Next == nil {
// 		return head
// 	}
// 	if head.Val != head.Next.Val {
// 		head.Next = deleteDuplicates(head.Next)
// 		return head
// 	} else {
// 		cur := head
// 		for cur.Next != nil && cur.Val == cur.Next.Val {
// 			cur = cur.Next
// 		}
// 		return deleteDuplicates(cur.Next)
// 	}
// }

// @lc code=end

