package bfs

import (
	"strconv"
)

/*
Key Points

1. Handling the problem with string

* TC

1. Worst to check all the nodes O(N), N = 10^4 = 10000

2. In each node, need to
  - check if current equal to target, 4
  - get the neighbors, 4 * 2 * (string handling) = 8
  - check if the neighbors in the dead map and visited map. 2 * 8 * (string comparing, around 4) = 64
  - totally is 64 + 8 + 4 = 76

3. Over all, it's about O(76*N) = O(N)

* SC

1. Queue, O(N), N = 10^4 = 10000
2. 2 Maps. O(N+L), L = len(deadends)
3. For each node, need to store the neighbors, around 8
4. Over all, it's O(N*8+N+L) = O(N)
*/

func openLock(deadends []string, target string) int {

	result := 0
	dMap := make(map[string]struct{})
	visited := make(map[string]struct{})
	for _, str := range deadends {
		dMap[str] = struct{}{}
	}

	q := []string{"0000"}

	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {
			current := q[idx]

			if _, exists := dMap[current]; exists {
				continue
			}
			if current == target {
				return result
			}
			neighbors := getNeighbors(current)
			for _, neighbor := range neighbors {
				if _, exists := dMap[neighbor]; exists {
					continue
				}
				if _, exists := visited[neighbor]; exists {
					continue
				}
				q = append(q, neighbor)
				visited[neighbor] = struct{}{}
			}

		}
		result++
		q = q[l:]
	}

	return -1
}

func getNeighbors(current string) []string {
	r := []string{}
	b := []byte(current)
	for idx, d := range b {
		if d == byte('9') {
			b[idx] = '0'
			r = append(r, string(b))
			b[idx] = '8'
			r = append(r, string(b))
			b[idx] = '9'
			continue
		}
		if d == byte('0') {
			b[idx] = '1'
			r = append(r, string(b))
			b[idx] = '9'
			r = append(r, string(b))
			b[idx] = '0'
			continue
		}
		b[idx] += 1
		r = append(r, string(b))
		b[idx] -= 2
		r = append(r, string(b))
		b[idx] += 1
	}

	return r
}

/*
Key Points:

1. Optimized version. Handling with integer. It would be much faster than that of string

* TC:

1. Check all node O(N). N = 10^4

 2. At each node, need to
    a. Get all the neighbors. 4 * 2 * (integer handling)
    b. Check if the neighbor in visited. 8 * (integer comparing, around 1)

3. Over all, it's round O(N*16) = O(N)

* SC

1. Queue. O(N)
2. Visited map O(N)
3. On each node checking, neighbors as O(8)
4. Over all, around O(8N+N) = O(N)
*/
func openLockOptimized(deadends []string, target string) int {

	result := 0
	targetInt, _ := strconv.Atoi(target)
	visited := [10000]int{}
	for _, str := range deadends {
		n, _ := strconv.Atoi(str)
		visited[n] = 1
	}

	q := []int{0}

	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {
			current := q[idx]

			if visited[current] == 1 {
				continue
			}
			if current == targetInt {
				return result
			}

			visited[current] = 1
			neighbors := getNeighborsInt(current)
			for _, neighbor := range neighbors {
				if visited[neighbor] == 1 {
					continue
				}
				q = append(q, neighbor)
			}
		}
		result++
		q = q[l:]
	}

	return -1
}

func getNeighborsInt(current int) [8]int {

	result := [8]int{}
	changeFactors := [4]int{1000, 100, 10, 1}

	for i, factor := range changeFactors {
		d := (current / factor) % 10
		for j, c := range [2]int{-1, 1} {
			z := d + c
			z = (z + 10) % 10
			result[i*2+j] = current + (z-d)*factor
		}
	}

	return result
}
