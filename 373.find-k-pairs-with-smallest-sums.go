/*
 * @lc app=leetcode id=373 lang=golang
 *
 * [373] Find K Pairs with Smallest Sums
 */

// @lc code=start
import (
	"container/heap"
)

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := PairHeap{}
	cnt := 0
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			heap.Push(&h, Pair{nums1[i], nums2[j]})
			cnt++
		}
	}
	res := make([][]int, 0)
	if k > cnt {
		k = cnt
	}
	for i := 0; i < k; i++ {
		pair := heap.Pop(&h).(Pair)
		res = append(res, pair)
	}
	return res
}

type Pair = []int

type PairHeap struct {
	list []Pair
}

func (ph PairHeap) Len() int { return len(ph.list) }
func (ph PairHeap) Less(i, j int) bool {
	return (ph.list[i][0] + ph.list[i][1]) < (ph.list[j][0] + ph.list[j][1])
}
func (ph PairHeap) Swap(i, j int) {
	ph.list[i], ph.list[j] = ph.list[j], ph.list[i]
}
func (ph *PairHeap) Push(x interface{}) {
	ph.list = append(ph.list, x.(Pair))
}
func (ph *PairHeap) Pop() interface{} {
	pair := ph.list[len(ph.list)-1]
	ph.list = ph.list[:len(ph.list)-1]
	return pair
}

// @lc code=end

