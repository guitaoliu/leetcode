/*
 * @lc app=leetcode id=165 lang=golang
 *
 * [165] Compare Version Numbers
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	version11 := strings.Split(version1, ".")
	version22 := strings.Split(version2, ".")
	i := 0
	for ; i < min(len(version11), len(version22)); i++ {
		v1, _ := strconv.Atoi(version11[i])
		v2, _ := strconv.Atoi(version22[i])
		if v1 > v2 {
			return 1
		} else if v1 < v2 {
			return -1
		} else {
			continue
		}
	}
	if len(version11) > len(version22) {
		for j := i; j < len(version11); j++ {
			v, _ := strconv.Atoi(version11[j])
			if v > 0 {
				return 1
			} else {
				continue
			}
		}
	} else if len(version22) > len(version11) {
		for j := i; j < len(version22); j++ {
			v, _ := strconv.Atoi(version22[j])
			if v > 0 {
				return -1
			} else {
				continue
			}
		}
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

