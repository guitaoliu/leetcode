/*
 * @lc app=leetcode id=739 lang=golang
 *
 * [739] Daily Temperatures
 */

// @lc code=start
func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	toCheck := []int{}
	for i, t := range temperatures {
		for len(toCheck) > 0 && temperatures[toCheck[len(toCheck)-1]] < t {
			idx := toCheck[len(toCheck)-1]
			toCheck = toCheck[:len(toCheck)-1]
			res[idx] = i - idx
		}
		toCheck = append(toCheck, i)
	}
	return res
}

// @lc code=end

