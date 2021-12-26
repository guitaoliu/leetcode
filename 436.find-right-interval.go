/*
 * @lc app=leetcode id=436 lang=golang
 *
 * [436] Find Right Interval
 */

// @lc code=start
func findRightInterval(intervals [][]int) []int {
	kv := make([][]int, len(intervals))
	for i, interval := range intervals {
		kv[i] = []int{interval[0], i}
	}
	sort.Slice(kv, func(i, j int) bool { return kv[i][0] < kv[j][0] })

	find := func(target int) int {
		left, right := 0, len(kv)-1
		for left < right {
			mid := left + (right-left)>>1
			if kv[mid][0] == target {
				return kv[mid][1]
			} else if kv[mid][0] > target {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if kv[left][0] >= target {
			return kv[left][1]
		}
		return -1
	}

	ans := []int{}
	for _, interval := range intervals {
		loc := find(interval[1])
		ans = append(ans, loc)
	}

	return ans
}

func findRightIntervalHeap(intervals [][]int) []int {
    maxStartHeap := &MaxHeap{}
    maxEndHeap := &MaxHeap{}
    for i, interval := range intervals {
        heap.Push(maxStartHeap, &ValIndex{interval[0], i})
        heap.Push(maxEndHeap, &ValIndex{interval[1], i})
    }
    res := make([]int, len(intervals))
    for i := 0; i < len(intervals); i++ {
        maxEndValIdx := heap.Pop(maxEndHeap).(*ValIndex)
        res[maxEndValIdx.Idx] = -1
        if (*maxStartHeap)[0].Val >= maxEndValIdx.Val {
            maxStartValIndex := heap.Pop(maxStartHeap).(*ValIndex)
            for maxStartHeap.Len() > 0 && (*maxStartHeap)[0].Val >= maxEndValIdx.Val {
                maxStartValIndex = heap.Pop(maxStartHeap).(*ValIndex)
            }
            res[maxEndValIdx.Idx] = maxStartValIndex.Idx
            heap.Push(maxStartHeap, maxStartValIndex)
        }
    }
    return res
}

type ValIndex struct {
    Val int
    Idx int
}

type MaxHeap []*ValIndex

func (mh MaxHeap) Len() int           { return len(mh) }
func (mh MaxHeap) Less(i, j int) bool { return mh[i].Val > mh[j].Val }
func (mh MaxHeap) Swap(i, j int)      { mh[i], mh[j] = mh[j], mh[i] }

func (mh *MaxHeap) Pop() interface{} {
    old := *mh
    n := len(old)
    x := old[n-1]
    *mh = old[:n-1]
    return x
}

func (mh *MaxHeap) Push(x interface{}) {
    *mh = append(*mh, x.(*ValIndex))
}

// @lc code=end

