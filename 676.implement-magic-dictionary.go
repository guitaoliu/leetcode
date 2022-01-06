/*
 * @lc app=leetcode id=676 lang=golang
 *
 * [676] Implement Magic Dictionary
 */

// @lc code=start
type MagicDictionary struct {
	m map[int]string
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		m: make(map[int]string),
	}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for k, v := range dictionary {

	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	for _, word := range this.m {
		n := 0
		if len(word) == len(searchWord) {
			for i := 0; i < len(word); i++ {
				if word[i] != searchWord[i] {
					n++
				}
			}
		}
		if n == 1 {
			return true
		}
	}
	return false
}

/**
 * Your MagicDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.BuildDict(dictionary);
 * param_2 := obj.Search(searchWord);
 */
// @lc code=end

