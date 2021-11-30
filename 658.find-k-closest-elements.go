/*
 * @lc app=leetcode id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	if x <= arr[0] {
		return arr[:k]
	}
	if x >= arr[len(arr)-1] {
		return arr[len(arr)-k:]
	}
	idx := sort.SearchInts(arr, x)
	i, j := idx-1, idx
	for j-i-1 < k {
		if i == -1 {
			j++
			continue
		}
		if j == len(arr) {
			i--
			continue
		}
		left := abs(arr[i] - x)
		right := abs(arr[j] - x)
		if left < right {
			i--
		} else if left > right {
			j++
		} else {
			if arr[i] < arr[j] {
				i--
			} else if arr[i] > arr[j] {
				j++
			} else {
				i--
				j++
			}
		}
	}
	return arr[i+1 : j]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end

