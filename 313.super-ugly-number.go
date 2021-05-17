/*
 * @lc app=leetcode id=313 lang=golang
 *
 * [313] Super Ugly Number
 */

// @lc code=start
import (
	"container/heap"
	"sort"
)

type MinHeap struct {
	sort.IntSlice
}

func (h *MinHeap) Push(x interface{}) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	a := h.IntSlice
	res := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return res
}

func nthSuperUglyNumber(n int, primes []int) int {
	h := &MinHeap{sort.IntSlice{1}}
	heap.Init(h)
	m := map[int]struct{}{1: struct{}{}}
	for i := 1; i < n; i++ {
		x := heap.Pop(h).(int)
		for _, f := range primes {
			next := x * f
			if _, ok := m[next]; !ok {
				heap.Push(h, next)
				m[next] = struct{}{}
			}
		}
	}
	return heap.Pop(h).(int)
}

// @lc code=end

