/*
 * @lc app=leetcode id=641 lang=golang
 *
 * [641] Design Circular Deque
 */

// @lc code=start
type ListNode struct {
	Val  int
	Prev *ListNode
	Next *ListNode
}

type MyCircularDeque struct {
	head *ListNode
	tail *ListNode
	l    int
	k    int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{k: k}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue = append([]int{value}, this.queue...)
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue = append(this.queue, value)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.queue = this.queue[1:]
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.queue = this.queue[:len(this.queue)-1]
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[0]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[len(this.queue)-1]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return len(this.queue) == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return len(this.queue) == this.k
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

