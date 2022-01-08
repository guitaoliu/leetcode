/*
 * @lc app=leetcode id=692 lang=golang
 *
 * [692] Top K Frequent Words
 */

// @lc code=start
type WordCnt struct {
	word string
	cnt  int
}
type FreqHeap []*WordCnt

func (a FreqHeap) Len() int      { return len(a) }
func (a FreqHeap) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a FreqHeap) Less(i, j int) bool {
	if a[i].cnt > a[j].cnt {
		return true
	} else if a[i].cnt == a[j].cnt {
		return a[i].word < a[j].word
	} else {
		return false
	}
}

func (a *FreqHeap) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[:n-1]
	return x
}
func (a *FreqHeap) Push(x interface{}) {
	*a = append(*a, x.(*WordCnt))
}

func topKFrequent(words []string, k int) []string {
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}

	h := &FreqHeap{}

	for word, cnt := range m {
		heap.Push(h, &WordCnt{word, cnt})
	}
	res := []string{}
	for i := 0; i < k; i++ {
		w := heap.Pop(h).(*WordCnt)
		res = append(res, w.word)
	}
	return res
}

// @lc code=end

