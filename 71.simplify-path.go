/*
 * @lc app=leetcode id=71 lang=golang
 *
 * [71] Simplify Path
 */

// @lc code=start
func simplifyPath(path string) string {
	strs := strings.Split(path, "/")
	res := make([]string, 0)
	for _, str := range strs {
		switch str {
		case "", ".":
			continue
		case "..":
			if len(res) > 0 {
				res = res[:len(res)-1]
			}
		default:
			res = append(res, str)
		}
	}
	return "/" + strings.Join(res, "/")
}

// @lc code=end

