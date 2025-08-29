package bfs

/*
KP

* TC
1. The top bound of nodes to check should be O(N) = 4^8, but there at most O(L), L is length of bank
2. But for each node, we need to generate the neighbors of it, that's 3^8
3. So, over all, it's O(L*3^8)

* SC
1. Queue size O(L)
2. Bank Map O(L)
3. Visited Map O(L)
4. At each node, there maybe at most L valid mutations. O(L). This part can be optimized by building the mutations in-line
5. Over all, is O(L^2)
*/

func minMutation(startGene string, endGene string, bank []string) int {

	result := 0
	visited := make(map[string]struct{})
	validMap := make(map[string]struct{})
	for _, b := range bank {
		validMap[b] = struct{}{}
	}

	visited[startGene] = struct{}{}
	q := []string{startGene}

	for len(q) > 0 {

		l := len(q)

		for idx := 0; idx < l; idx++ {
			current := q[idx]
			if current == endGene {
				return result
			}
			mutations := generateValidMutation(current, validMap) // can be optimized if we build this in-line
			mutationsToProcess := []string{}
			for _, gene := range mutations {
				if _, exist := visited[gene]; exist {
					continue
				}
				mutationsToProcess = append(mutationsToProcess, gene)
				visited[gene] = struct{}{}
			}
			if len(mutationsToProcess) > 0 {
				result++
				q = append(q, mutationsToProcess...)
			}
		}

		q = q[l:]
	}

	return -1
}

func generateValidMutation(start string, validMap map[string]struct{}) []string {

	changeMap := map[rune][]rune{
		'A': {'C', 'G', 'T'},
		'C': {'A', 'G', 'T'},
		'G': {'A', 'C', 'T'},
		'T': {'A', 'C', 'G'},
	}
	result := []string{}

	for idx := 0; idx < 8; idx++ {
		for _, change := range changeMap[rune(start[idx])] {
			mutation := start[:idx] + string(change) + start[idx+1:]
			if _, exist := validMap[mutation]; !exist {
				continue
			}
			result = append(result, mutation)
		}
	}
	return result
}
