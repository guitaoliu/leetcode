/*
 * @lc app=leetcode id=284 lang=golang
 *
 * [284] Peeking Iterator
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iter     *Iterator
	nextElem int
	hasElem  bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iter: iter}
}

func (this *PeekingIterator) hasNext() bool {
	if this.hasElem {
		return true
	}
	return this.iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.hasElem {
		this.hasElem = false
		return this.nextElem
	}
	return this.iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.hasElem {
		return this.nextElem
	}
	this.hasElem = true
	this.nextElem = this.iter.next()
	return this.nextElem
}

// @lc code=end

