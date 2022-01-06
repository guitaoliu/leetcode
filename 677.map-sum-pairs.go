/*
 * @lc app=leetcode id=677 lang=golang
 *
 * [677] Map Sum Pairs
 */

// @lc code=start
type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{make(map[string]int)}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	cnt := 0
	for word, val := range this.m {
		if strings.HasPrefix(word, prefix) {
			cnt += val
		}
	}
	return cnt
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
// @lc code=end

