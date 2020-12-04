/*
 * @lc app=leetcode id=506 lang=golang
 *
 * [506] Relative Ranks
 */

// @lc code=start
import (
	"sort"
	"strconv"
)

func findRelativeRanks(nums []int) []string {
	gradeMap := make(map[int]string, len(nums))
	sortedNums := make([]int, len(nums))
	copy(sortedNums, nums)
	res := make([]string, len(nums))
	sort.Sort(sort.Reverse(sort.IntSlice(sortedNums)))
	for idx, v := range sortedNums {
		gradeMap[v] = getGrade(idx + 1)
	}
	for idx, v := range nums {
		res[idx] = gradeMap[v]
	}
	return res
}

func getGrade(num int) string {
	switch num {
	case 1:
		return "Gold Medal"
	case 2:
		return "Silver Medal"
	case 3:
		return "Bronze Medal"
	default:
		return strconv.Itoa(num)
	}
}

// @lc code=end

