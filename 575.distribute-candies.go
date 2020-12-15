/*
 * @lc app=leetcode id=575 lang=golang
 *
 * [575] Distribute Candies
 */

// @lc code=start
func distributeCandies(candyType []int) int {
	identicalCandy := make(map[int]bool, len(candyType))
	cnt := 0
	for _, v := range candyType {
		if _, ok := identicalCandy[v]; !ok {
			identicalCandy[v] = true
			cnt++
		}
	}
	return min(cnt, len(candyType)/2)
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

// @lc code=end

