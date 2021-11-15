/*
 * @lc app=leetcode id=622 lang=golang
 *
 * [622] Design Circular Queue
 */

// @lc code=start
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}
type MyCircularQueue struct {
	Head     *Node
	Tail     *Node
	Length   int
	Capacity int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{Capacity: k}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	node := &Node{Val: value}
	if this.IsEmpty() {
		this.Head = node
		this.Tail = node
	}
	node.Prev = this.Tail
	this.Tail.Next = node
	this.Tail = node
	this.Length++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	if this.Length == 1 {
		this.Head = nil
		this.Tail = nil
	} else {
		head := this.Head
		this.Head = this.Head.Next
		this.Head.Prev = nil
		head.Next = nil
	}
	this.Length--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.Head == nil {
		return -1
	}
	return this.Head.Val
}

func (this *MyCircularQueue) Rear() int {
	if this.Tail == nil {
		return -1
	}
	return this.Tail.Val
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.Length == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.Length == this.Capacity
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */
// @lc code=end

