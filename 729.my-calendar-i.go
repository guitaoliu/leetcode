/*
 * @lc app=leetcode id=729 lang=golang
 *
 * [729] My Calendar I
 */

// @lc code=start
type MyCalendar struct {
	reservations [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (this *MyCalendar) Book(start int, end int) bool {
	sPos := sort.Search(len(this.reservations), func(i int) bool { return start < this.reservations[i][1] })
	if sPos == len(this.reservations) {
		this.reservations = append(this.reservations, [2]int{start, end})
		return true
	}
	if this.reservations[sPos][0] >= end {
		this.reservations = append(this.reservations, [2]int{})
		copy(this.reservations[sPos+1:], this.reservations[sPos:])
		this.reservations[sPos] = [2]int{start, end}
		return true
	}
	return false
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end

