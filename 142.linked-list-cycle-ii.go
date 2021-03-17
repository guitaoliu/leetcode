/*
 * @lc app=leetcode id=142 lang=golang
 *
 * [142] Linked List Cycle II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}

// func detectCycle(head *ListNode) *ListNode {
// 	nodeMap := make(map[*ListNode]struct{}, 0)
// 	for head != nil {
// 		if _, ok := nodeMap[head]; ok {
// 			return head
// 		}
// 		nodeMap[head] = struct{}{}
// 		head = head.Next
// 	}
// 	return nil
// }

// @lc code=end

