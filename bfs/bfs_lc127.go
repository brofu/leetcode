package bfs

/*
KP

# Complexity

Let
L, the length of `beginWord`
N, the number in wordList

TC
1. Construct wordMap, O(N)
2. Nodes in the graph, around O(N), for each node, need to check O(25*L) sides.
3. Over all O(N*25*L), it's O(N*L)

SC:
1. Visited, wordMap around O(2N)
2. Queue, O(N)
3. Overall O(3N). O(N). If we account the word length in the space, then, it's O(N*L)
*/
func ladderLength(beginWord string, endWord string, wordList []string) int {

	visited := make(map[string]struct{})
	wordMap := make(map[string]struct{})
	for _, word := range wordList {
		wordMap[word] = struct{}{}
	}
	if _, exists := wordMap[endWord]; !exists {
		return 0
	}

	getNeighbors := func(word string) []string {
		result := []string{}
		for idx, letter := range word {
			for i := 'a'; i <= 'z'; i++ {
				if letter == i {
					continue
				}
				w := string(word[:idx]) + string(i) + string(word[idx+1:])
				if _, exists := wordMap[w]; !exists {
					continue
				}
				result = append(result, w)
			}
		}
		return result
	}

	result := 0
	visited[beginWord] = struct{}{}
	q := []string{beginWord}
	for len(q) > 0 {
		l := len(q)
		result++
		for idx := 0; idx < l; idx++ {
			current := q[idx]
			if current == endWord {
				return result
			}
			neighbors := getNeighbors(current)
			for _, neighbor := range neighbors {
				if _, exists := visited[neighbor]; exists {
					continue
				}
				visited[neighbor] = struct{}{}
				q = append(q, neighbor)
			}
		}
		q = q[l:]
	}

	return 0
}

func ladderLengthV2(beginWord string, endWord string, wordList []string) int {

	differList := make([][]int, len(beginWord)+1)
	endExisted := false

	for idx, word := range wordList {
		if word == endWord {
			endExisted = true
		}
		differNumber := differ(beginWord, word)
		differList[differNumber] = append(differList[differNumber], idx)
	}

	if !endExisted {
		return 0
	}

	return 0
}

func differ(s, d string) (differ int) {
	for i := 0; i < len(s); i++ {
		if s[i] != d[i] {
			differ++
		}
	}
	return
}
