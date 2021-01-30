/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	record := map[string][]string{}
	res := [][]string{}
	for _, str := range strs {
		sBytes := []byte(str)
		sort.Slice(sBytes, func(i, j int) bool { return sBytes[i] < sBytes[j] })
		sortedStr := string(sBytes)
		record[sortedStr] = append(record[sortedStr], str)
	}
	for _, v := range record {
		res = append(res, v)
	}
	return res
}

// @lc code=end

