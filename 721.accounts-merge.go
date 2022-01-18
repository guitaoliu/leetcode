/*
 * @lc app=leetcode id=721 lang=golang
 *
 * [721] Accounts Merge
 */

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	parents := make([]int, len(accounts))
	for i := range parents {
		parents[i] = i
	}
	var find func(int) int
	find = func(i int) int {
		if parents[i] != i {
			return find(parents[i])
		}
		return i
	}

	union := func(a, b int) {
		x, y := find(a), find(b)
		parents[x] = y
	}
	idToName := make(map[int]string)
	emailToId := make(map[string]int)

	for id := 0; id < len(accounts); id++ {
		acc := accounts[id]
		idToName[id] = acc[0]
		for i := 1; i < len(acc); i++ {
			email := acc[i]
			if pid, ok := emailToId[email]; ok {
				union(id, pid)
			}
			emailToId[email] = id
		}
	}

	idToEmails := make(map[int][]string)
	for email, id := range emailToId {
		pid := find(id)
		idToEmails[pid] = append(idToEmails[pid], email)
	}

	res := [][]string{}
	for id, emails := range idToEmails {
		name := idToName[id]
		sort.Strings(emails)
		res = append(res, append([]string{name}, emails...))
	}

	return res
}

// @lc code=end

