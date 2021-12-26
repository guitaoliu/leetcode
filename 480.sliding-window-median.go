/*
 * @lc app=leetcode id=480 lang=golang
 *
 * [480] Sliding Window Median
 */

// @lc code=start
import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type MedianHeap struct {
	minHeap *IntHeap
	maxHeap *IntHeap
}

func (mh *MedianHeap) AddNum(num int) {
	if mh.maxHeap.Len() == 0 || num > (*mh.maxHeap)[0] {
		heap.Push(mh.maxHeap, num)
	} else {
		heap.Push(mh.minHeap, -num)
	}
	mh.Reblance()
}

func (mh *MedianHeap) Reblance() {
	if mh.maxHeap.Len() > mh.minHeap.Len()+1 {
		heap.Push(mh.minHeap, -heap.Pop(mh.maxHeap).(int))
	}
	if mh.minHeap.Len() > mh.maxHeap.Len() {
		heap.Push(mh.maxHeap, -heap.Pop(mh.minHeap).(int))
	}
}

func (mh *MedianHeap) FindMedian() float64 {
	if mh.minHeap.Len() == mh.maxHeap.Len() {
		return float64(-(*mh.minHeap)[0]+(*mh.maxHeap)[0]) / 2
	} else {
		return float64((*mh.maxHeap)[0])
	}
}

func (mh *MedianHeap) Remove(num int) {
	if num >= (*mh.maxHeap)[0] {
		for i, n := range *mh.maxHeap {
			if n == num {
				heap.Remove(mh.maxHeap, i)
				break
			}
		}
	} else {
		for i, n := range *mh.minHeap {
			if n == -num {
				heap.Remove(mh.minHeap, i)
				break
			}
		}
	}
	mh.Reblance()
}

func medianSlidingWindow(nums []int, k int) []float64 {
	res := make([]float64, len(nums)-k+1)
	p := 0
	mh := &MedianHeap{
		maxHeap: &IntHeap{},
		minHeap: &IntHeap{},
	}
	for i, num := range nums {
		mh.AddNum(num)
		if i-k+1 >= 0 {
			res[p] = mh.FindMedian()
			mh.Remove(nums[i-k+1])
			p++
		}
	}
	return res
}

// @lc code=end

