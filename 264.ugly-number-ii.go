/*
 * @lc app=leetcode id=264 lang=golang
 *
 * [264] Ugly Number II
 */

// @lc code=start

func nthUglyNumber(n int) int {
	dp := make([]int, n+1)
	dp[1] = 1
	p2, p3, p5 := 1, 1, 1
	for i := 2; i <= n; i++ {
		x2, x3, x5 := dp[p2]*2, dp[p3]*3, dp[p5]*5
		dp[i] = min(x2, x3, x5)
		if dp[i] == x2 {
			p2++
		}
		if dp[i] == x3 {
			p3++
		}
		if dp[i] == x5 {
			p5++
		}
	}
	return dp[n]
}

func min(a, b, c int) int {
	minimum := a
	if b < minimum {
		minimum = b
	}
	if c < minimum {
		minimum = c
	}
	return minimum
}

// import (
// 	"container/heap"
// )

// type MinHeap []int

// func (h MinHeap) Len() int            { return len(h) }
// func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
// func (h *MinHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
// func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
// func (h *MinHeap) Pop() interface{} {
// 	res := (*h)[len(*h)-1]
// 	*h = (*h)[:len(*h)-1]
// 	return res
// }

// var factors = []int{2, 3, 5}

// func nthUglyNumber(n int) int {
// 	h := &MinHeap{1}
// 	heap.Init(h)
// 	m := map[int]struct{}{1: struct{}{}}
// 	for i := 1; ; i++ {
// 		x := heap.Pop(h).(int)
// 		if i == n {
// 			return x
// 		}
// 		for _, f := range factors {
// 			next := x * f
// 			if _, ok := m[next]; !ok {
// 				heap.Push(h, next)
// 				m[next] = struct{}{}
// 			}
// 		}
// 	}
// }

// @lc code=end

