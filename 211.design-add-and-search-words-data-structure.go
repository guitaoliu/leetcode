/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */

// @lc code=start
type WordDictionary struct {
	children map[rune]*WordDictionary
	isWord   bool
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{children: make(map[rune]*WordDictionary)}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, char := range word {
		if child, ok := cur.children[char]; ok {
			cur = child
		} else {
			newChild := &WordDictionary{children: make(map[rune]*WordDictionary)}
			cur.children[char] = newChild
			cur = newChild
		}
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	cur := this
	for i, ch := range word {
		if ch == '.' {
			isMatched := false
			for _, v := range cur.children {
				if v.Search(word[i+1:]) {
					isMatched = true
				}
			}
			return isMatched
		} else if child, ok := cur.children[ch]; ok {
			cur = child
		} else {
			return false
		}
	}
	return cur.isWord == true
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

