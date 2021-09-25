/*
 * @lc app=leetcode id=445 lang=golang
 *
 * [445] Add Two Numbers II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	len1, len2 := getLen(l1), getLen(l2)
	if len1 < len2 {
		l1, l2 = l2, l1
		len1, len2 = len2, len1
	}
	newHead := &ListNode{Next: l1}
	newHead.Next = add(l1, l2, len1-len2)

	if newHead.Next != nil && newHead.Next.Val >= 10 {
		newHead.Next.Val = newHead.Next.Val % 10
		newHead.Val++
		return newHead
	}
	return newHead.Next
}

func add(l1 *ListNode, l2 *ListNode, leading int) *ListNode {
	if l1 == nil {
		return nil
	}
	var node *ListNode
	if leading > 0 {
		node = &ListNode{Val: l1.Val}
		node.Next = add(l1.Next, l2, leading-1)
	} else {
		node = &ListNode{Val: l1.Val + l2.Val}
		node.Next = add(l1.Next, l2.Next, 0)
	}
	if node.Next != nil && node.Next.Val >= 10 {
		node.Next.Val = node.Next.Val % 10
		node.Val++
	}
	return node
}

func getLen(node *ListNode) int {
	cnt := 0
	for node != nil {
		cnt++
		node = node.Next
	}
	return cnt
}

// @lc code=end

