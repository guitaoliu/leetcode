/*
 * @lc app=leetcode id=341 lang=golang
 *
 * [341] Flatten Nested List Iterator
 */

// @lc code=start
/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */
type NestedIterator struct {
	stack [][]*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{[][]*NestedInteger{nestedList}}
}

func (this *NestedIterator) Next() int {
	res := this.stack[len(this.stack)-1][0]
	this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
	return res.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stack) > 0 {
		queue := this.stack[len(this.stack)-1]
		if len(queue) == 0 {
			this.stack = this.stack[:len(this.stack)-1]
			continue
		}
		if queue[0].IsInteger() {
			return true
		}
		newList := queue[0].GetList()
		this.stack[len(this.stack)-1] = this.stack[len(this.stack)-1][1:]
		this.stack = append(this.stack, newList)
	}
	return false
}

// type NestedIterator struct {
// 	ptr  int
// 	list []int
// }

// func Constructor(nestedList []*NestedInteger) *NestedIterator {
// 	list := construct(nestedList)
// 	return &NestedIterator{list: list}
// }

// func (this *NestedIterator) Next() int {
// 	cur := this.ptr
// 	this.ptr++
// 	return this.list[cur]
// }

// func (this *NestedIterator) HasNext() bool {
// 	return this.ptr < len(this.list)
// }

// func construct(nestedList []*NestedInteger) []int {
// 	list := []int{}
// 	for _, item := range nestedList {
// 		if item.IsInteger() {
// 			list = append(list, item.GetInteger())
// 		} else {
// 			flattenList := construct(item.GetList())
// 			list = append(list, flattenList...)
// 		}
// 	}
// 	return list
// }

// @lc code=end

