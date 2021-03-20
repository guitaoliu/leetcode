/*
 * @lc app=leetcode id=146 lang=golang
 *
 * [146] LRU Cache
 */

// @lc code=start
type LRUCache struct {
	head, tail *Node
	Keys       map[int]*Node
	cap        int
}

type Node struct {
	Prev, Next *Node
	Key, Val   int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{Keys: make(map[int]*Node), cap: capacity}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Keys[key]; ok {
		this.remove(node)
		this.add(node)
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Keys[key]; ok {
		node.Val = value
		this.remove(node)
		this.add(node)
		return
	} else {
		node = &Node{Key: key, Val: value}
		this.Keys[key] = node
		this.add(node)
	}
	if len(this.Keys) > this.cap {
		delete(this.Keys, this.tail.Key)
		this.remove(this.tail)
	}
}

func (this *LRUCache) add(node *Node) {
	node.Next = this.head
	node.Prev = nil
	if this.head != nil {
		this.head.Prev = node
	}
	this.head = node
	if this.tail == nil {
		this.tail = node
	}
}

func (this *LRUCache) remove(node *Node) {
	if node == this.head {
		this.head = this.head.Next
		node.Next = nil
		return
	}
	if node == this.tail {
		this.tail = node.Prev
		node.Prev.Next = nil
		node.Prev = nil
		return
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end

