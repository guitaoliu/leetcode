/*
 * @lc app=leetcode id=735 lang=golang
 *
 * [735] Asteroid Collision
 */

// @lc code=start
func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	for _, asteroid := range asteroids {
		if len(stack) == 0 || asteroid*stack[len(stack)-1] > 0 {
			stack = append(stack, asteroid)
			continue
		}
		for len(stack) > 0 && (asteroid < 0 && stack[len(stack)-1] > 0) {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			asteroid = compare(asteroid, last)
			if asteroid == 0 {
				break
			}
		}
		if asteroid != 0 {
			stack = append(stack, asteroid)
		}
	}
	return stack
}

func compare(asteroid, last int) int {
	if -asteroid > last {
		return asteroid
	} else if -asteroid < last {
		return last
	} else {
		return 0
	}
}

// @lc code=end

