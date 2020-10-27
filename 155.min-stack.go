/*
 * @lc app=leetcode id=155 lang=golang
 *
 * [155] Min Stack
 */

// @lc code=start
type MinStack struct {
	elements []int
	min      []int
	length   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0), 0}
}

func (this *MinStack) Push(x int) {
	this.elements = append(this.elements, x)
	if this.length == 0 {
		this.min = append(this.min, x)
	} else {
		min := this.GetMin()
		if x < min {
			this.min = append(this.min, x)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.length++
}

func (this *MinStack) Pop() {
	this.length--
	this.min = this.min[:this.length]
	this.elements = this.elements[:this.length]
}

func (this *MinStack) Top() int {
	return this.elements[this.length-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.length-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

