/*
 * @lc app=leetcode id=378 lang=golang
 *
 * [378] Kth Smallest Element in a Sorted Matrix
 */

// @lc code=start
import (
	"container/heap"
	"sort"
)

func kthSmallestHeap(matrix [][]int, k int) int {
	h := &MinHeap{}
	for i := 0; i < len(matrix); i++ {
		heap.Push(h, [3]int{matrix[i][0], i, 0})
	}
	for i := 0; i < k-1; i++ {
		cur := heap.Pop(h).([3]int)
		if cur[2] != len(matrix)-1 {
			heap.Push(h, [3]int{matrix[cur[1]][cur[2]+1], cur[1], cur[2] + 1})
		}
	}
	return heap.Pop(h).([3]int)[0]
}

type MinHeap [][3]int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.([3]int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func kthSmallestSort(matrix [][]int, k int) int {
	slice := make([]int, len(matrix)*len(matrix[0]))
	idx := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			slice[idx] = matrix[i][j]
			idx++
		}
	}
	sort.Ints(slice)
	return slice[k-1]
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)>>1
		if check(matrix, mid, k, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(matrix [][]int, mid, k, n int) bool {
	i, j := n-1, 0
	num := 0
	for i >= 0 && j < n {
		if matrix[i][j] <= mid {
			num += i + 1
			j++
		} else {
			i--
		}
	}
	return num >= k
}

// @lc code=end

