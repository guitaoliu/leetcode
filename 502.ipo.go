/*
 * @lc app=leetcode id=502 lang=golang
 *
 * [502] IPO
 */

// @lc code=start
type NumIdx struct {
	val int
	idx int
}

type MinHeap []*NumIdx

func (mh MinHeap) Len() int           { return len(mh) }
func (mh MinHeap) Less(i, j int) bool { return mh[i].val < mh[j].val }
func (mh MinHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[:n-1]
	return x
}

func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(*NumIdx))
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	minCapital := &MinHeap{}
	maxProfit := &MinHeap{}

	for i, c := range capital {
		heap.Push(minCapital, &NumIdx{c, i})
	}

	for i := 0; i < k; i++ {
		for (*minCapital).Len() > 0 && (*minCapital)[0].val <= w {
			idx := heap.Pop(minCapital).(*NumIdx).idx
			heap.Push(maxProfit, &NumIdx{-profits[idx], idx})
		}

		if (*maxProfit).Len() == 0 {
			break
		}

		w += -heap.Pop(maxProfit).(*NumIdx).val
	}
	return w
}

// @lc code=end

