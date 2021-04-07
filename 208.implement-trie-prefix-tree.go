/*
 * @lc app=leetcode id=208 lang=golang
 *
 * [208] Implement Trie (Prefix Tree)
 */

// @lc code=start
type Trie struct {
	isEnd    bool
	children map[rune]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this
	for _, c := range word {
		if child, ok := cur.children[c]; ok {
			cur = child
		} else {
			newChild := &Trie{children: make(map[rune]*Trie)}
			cur.children[c] = newChild
			cur = newChild
		}
	}
	cur.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		if child, ok := cur.children[c]; ok {
			cur = child
		} else {
			return false
		}
	}
	return cur.isEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		if child, ok := cur.children[c]; ok {
			cur = child
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end

