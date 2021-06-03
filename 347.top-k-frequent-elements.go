/*
 * @lc app=leetcode id=347 lang=golang
 *
 * [347] Top K Frequent Elements
 */

// @lc code=start
import (
	"container/heap"
)

type Item struct {
	key   int
	count int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].count > pq[j].count }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}
func (pq *PriorityQueue) Pop() interface{} {
	item := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return item
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	q := PriorityQueue{}
	for key, val := range m {
		heap.Push(&q, &Item{key, val})
	}
	res := []int{}
	for i := 0; i < k; i++ {
		item := heap.Pop(&q).(*Item)
		res = append(res, item.key)
	}
	return res
}

// @lc code=end

