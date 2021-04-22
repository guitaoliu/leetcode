/*
 * @lc app=leetcode id=228 lang=golang
 *
 * [228] Summary Ranges
 */

// @lc code=start
func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	res := []string{}
	start, cur := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != cur+1 {
			s := strconv.Itoa(start)
			if start != cur {
				s += "-" + ">" + strconv.Itoa(cur)
			}
			res = append(res, s)
			start, cur = nums[i], nums[i]
		} else {
			cur = nums[i]
		}
		if i == len(nums)-1 {
			s := strconv.Itoa(start)
			if start != cur {
				s += "-" + ">" + strconv.Itoa(cur)
			}
			res = append(res, s)
		}
	}
	return res
}

// @lc code=end

