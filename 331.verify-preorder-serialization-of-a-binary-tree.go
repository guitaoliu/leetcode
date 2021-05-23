/*
 * @lc app=leetcode id=331 lang=golang
 *
 * [331] Verify Preorder Serialization of a Binary Tree
 */

// @lc code=start
import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	diff := 1
	for _, node := range nodes {
		diff--
		if diff < 0 {
			return false
		}
		if node != "#" {
			diff += 2
		}
	}
	return diff == 0
}

func isValidSerializationStack(preorder string) bool {
	nodes := strings.Split(preorder, ",")
	stack := []int{1}
	for _, node := range nodes {
		if len(stack) == 0 {
			return false
		}
		stack[len(stack)-1]--
		if stack[len(stack)-1] == 0 {
			stack = stack[:len(stack)-1]
		}
		if node != "#" {
			stack = append(stack, 2)
		}
	}
	return len(stack) == 0
}

// @lc code=end

