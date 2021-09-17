/*
 * @lc app=leetcode id=433 lang=golang
 *
 * [433] Minimum Genetic Mutation
 */

// @lc code=start
func minMutationBFS(start string, end string, bank []string) int {
	b := make(map[string]struct{})
	for _, g := range bank {
		b[g] = struct{}{}
	}

	genes := []byte{'A', 'C', 'G', 'T'}

	queue := []string{start}
	cnt := 0

	for len(queue) > 0 {
		curLevel := queue
		queue = nil

		for _, gene := range curLevel {
			if gene == end {
				return cnt
			}
			geneByte := []byte(gene)
			for i := 0; i < len(start); i++ {
				for _, v := range genes {
					if v == gene[i] {
						continue
					}
					newGene := make([]byte, len(gene))
					copy(newGene, geneByte)
					newGene[i] = v
					newGeneString := string(newGene)
					if _, ok := b[string(newGeneString)]; ok {
						queue = append(queue, newGeneString)
					}
				}
			}
		}
		cnt++
	}

	return -1
}

func minMutation(start, end string, bank []string) int {
	bankSet := make(map[string]struct{}, len(bank))
	for _, v := range bank {
		bankSet[v] = struct{}{}
	}

	if _, has := bankSet[end]; !has {
		return -1
	}

	visited := make(map[string]struct{})
	startSet, endSet := make(map[string]struct{}), make(map[string]struct{})
	startSet[start], endSet[end] = struct{}{}, struct{}{}

	step := 1
	geneBase := []byte{'A', 'C', 'G', 'T'}

	for len(startSet) > 0 && len(endSet) > 0 {
		if len(startSet) > len(endSet) {
			startSet, endSet = endSet, startSet
		}

		nextSet := make(map[string]struct{})
		for curr := range startSet {
			for i := 0; i < len(curr); i++ {
				for _, char := range geneBase {
					if curr[i] == char {
						continue
					}
					newGene := curr[:i] + string(char) + curr[i+1:]
					if _, ok := bankSet[newGene]; ok {
						if _, ok := endSet[newGene]; ok {
							return step
						}
						if _, ok := visited[newGene]; !ok {
							nextSet[newGene] = struct{}{}
							visited[newGene] = struct{}{}
						}
					}
				}
			}
		}

		startSet = nextSet
		step++
	}

	return -1
}

// @lc code=end

