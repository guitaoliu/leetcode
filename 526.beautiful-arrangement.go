/*
 * @lc app=leetcode id=526 lang=golang
 *
 * [526] Beautiful Arrangement
 */

// @lc code=start
import (
	"math/bits"
)

func countArrangementBrute(n int) int {
	res := []int{0, 1, 2, 3, 8, 10, 36, 41, 132, 250, 700, 750, 4010, 4237, 10680, 24679, 87328, 90478, 435812}
	return res[n]
}

func countArrangementDFS(n int) int {
	res := 0
	used := make(map[int]struct{})
	var trace func(int)
	trace = func(idx int) {
		if idx > n {
			res++
		} else {
			for i := 1; i <= n; i++ {
				if _, ok := used[i]; !ok {
					if i%idx == 0 || idx%i == 0 {
						used[i] = struct{}{}
						trace(idx + 1)
						delete(used, i)
					}
				}
			}
		}
	}
	trace(1)
	return res
}

func countArrangementCompressedMemory(n int) int {
	var trace func(int, int) int
	trace = func(i, used int) (cnt int) {
		if i > n {
			return 1
		} else {
			for num := 1; num <= n; num++ {
				if 1<<num&used == 0 && (num%i == 0 || i%num == 0) {
					cnt += trace(i+1, used|1<<num)
				}
			}
		}
		return
	}
	return trace(1, 0)
}

func countArrangementCompressedMemoryWithMemo(n int) int {
	m := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		m[i] = make([]int, 1<<n)
	}

	var trace func(int, int) int
	trace = func(i, used int) (cnt int) {
		if i > n {
			return 1
		}

		if m[i][used] != 0 {
			return m[i][used]
		}

		for num := 1; num <= n; num++ {
			if 1<<(num-1)&used == 0 && (num%i == 0 || i%num == 0) {
				cnt += trace(i+1, used|1<<(num-1))
			}
		}
		return
	}
	return trace(1, 0)
}

func countArrangementDP(n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 1<<n)
	}

	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for visited := 0; visited < 1<<n; visited++ {
			if bits.OnesCount(uint(visited)) == i {
				for num := 1; num <= n; num++ {
					if 1<<(num-1)&visited != 0 && (num%i == 0 || i%num == 0) {
						dp[i][visited] += dp[i-1][^(1<<(num-1))&visited]
					}
				}
			}
		}
	}
	return dp[n][1<<n-1]
}

func countArrangement(n int) int {
	dp := make([]int, 1<<n)
	dp[0] = 1

	for visited := 0; visited < 1<<n; visited++ {
		i := bits.OnesCount(uint(visited))
		for num := 1; num <= n; num++ {
			if 1<<(num-1)&visited != 0 && (num%i == 0 || i%num == 0) {
				dp[visited] += dp[visited & ^(1<<(num-1))]
			}
		}
	}
	return dp[1<<n-1]
}

// @lc code=end

