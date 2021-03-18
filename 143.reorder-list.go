/*
 * @lc app=leetcode id=143 lang=golang
 *
 * [143] Reorder List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow

	cur := mid.Next
	for cur.Next != nil {
		tmp := cur.Next
		cur.Next = tmp.Next
		tmp.Next = mid.Next
		mid.Next = tmp
	}

	l1 := head
	l2 := mid.Next
	mid.Next = nil
	for l1 != nil && l2 != nil {
		l1Tmp := l1.Next
		l2Tmp := l2.Next
		l1.Next = l2
		l1 = l1Tmp
		l2.Next = l1
		l2 = l2Tmp
	}
}

// func reorderList(head *ListNode) {
// 	l := []*ListNode{}
// 	cur := head
// 	for cur != nil {
// 		l = append(l, cur)
// 		cur = cur.Next
// 	}
// 	for i, j := 0, len(l)-1; i < j; i, j = i+1, j-1 {
// 		if l[i].Next == l[j] {
// 			l[j].Next = nil
// 			break
// 		}
// 		l[j].Next = l[i].Next
// 		l[i].Next = l[j]
// 	}
// 	if len(l)%2 != 0 {
// 		l[len(l)/2].Next = nil
// 	}
// }

// @lc code=end

