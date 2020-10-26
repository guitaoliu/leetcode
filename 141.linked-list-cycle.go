/*
 * @lc app=leetcode id=141 lang=golang
 *
 * [141] Linked List Cycle
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func hasCycle(head *ListNode) bool {
// 	visited := make(map[*ListNode]int)
// 	for head != nil {
// 		if _, ok := visited[head]; ok {
// 			return true
// 		}
// 		visited[head] = head.Val
// 		head = head.Next
// 	}
// 	return false
// }

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head, head
	for slow != nil && fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// @lc code=end

