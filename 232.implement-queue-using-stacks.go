/*
 * @lc app=leetcode id=232 lang=golang
 *
 * [232] Implement Queue using Stacks
 */

// @lc code=start
type MyQueue struct {
	Stack []int
	Queue []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{Stack: []int{}, Queue: []int{}}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	this.stackToQueue()
	popped := this.Queue[len(this.Queue)-1]
	this.Queue = this.Queue[:len(this.Queue)-1]
	return popped
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	this.stackToQueue()
	return this.Queue[len(this.Queue)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Stack)+len(this.Queue) == 0
}

func (this *MyQueue) stackToQueue() {
	if len(this.Queue) == 0 {
		for len(this.Stack) > 0 {
			this.Queue = append(this.Queue, this.Stack[len(this.Stack)-1])
			this.Stack = this.Stack[:len(this.Stack)-1]
		}
	}
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
// @lc code=end

