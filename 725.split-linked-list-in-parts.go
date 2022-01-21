/*
 * @lc app=leetcode id=725 lang=golang
 *
 * [725] Split Linked List in Parts
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func splitListToParts(head *ListNode, k int) []*ListNode {
	length := 0
	cur := head
	for cur != nil {
		length++
		cur = cur.Next
	}
	res := make([]*ListNode, k)
	n := length / k
	c := length % k
	cur = head
	for i := 0; i < k; i++ {
		var next *ListNode
		if i < c {
			next = getN(head, n+1)
		} else {
			next = getN(head, n)
		}
		res[i] = head
		head = next
	}
	return res
}

func getN(head *ListNode, n int) *ListNode {
	if n == 0 {
		return nil
	}
	cur := head
	for n-1 > 0 {
		cur = cur.Next
		n--
	}
	tmp := cur.Next
	cur.Next = nil
	return tmp
}

// @lc code=end

