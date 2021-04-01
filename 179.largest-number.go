/*
 * @lc app=leetcode id=179 lang=golang
 *
 * [179] Largest Number
 */

// @lc code=start

func largestNumber(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	strs := make([]string, len(nums))
	for _, v := range nums {
		str := strconv.Itoa(v)
		strs = append(strs, str)
	}
	sort.SliceStable(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	var res strings.Builder
	res.Grow(len(strs))
	for _, str := range strs {
		if str == "0" && res.String() == "0" {
			continue
		}
		res.WriteString(str)
	}
	return res.String()
}

// @lc code=end

