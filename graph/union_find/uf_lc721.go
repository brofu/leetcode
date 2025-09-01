package uf

import (
	"sort"
)

/*
KP
1. Faster to process with `Integer` than `String`
2. What's a(n)?
Ackermann A(m,n)	Very fast-growing recursive function
Inverse Ackermann α(n)	Grows incredibly slowly — practically ≤ 5
Use in CS	Time complexity of Union-Find with path compression

TC
Let
A = number of accounts
E = total number of email mentions (i.e., total number of entries in account[1:] across all accounts)
N = number of unique emails
1. Construct uf O(A)
2. Construct idxMap O(E). And for unionIndex, around a(n), so, O(E*a(n))
3. Construct resultMap. Around O(N*a(n))
4. Sort emails. O(N*lgN)
5. Overall, O(A + E*a(n) + N*a(n) + O(N*lgN)), around O(E*a(n)+N*lgN). (A <= E, N <= E), and a(n) is almost consistent

SC
1. uf, O(A)
2. idxMap O(N)
3. resultMap O(N)
4. result O(N)
5. Overall O(3N+A), around O(N+A)
*/

func accountsMerge(accounts [][]string) [][]string {

	uf := make([]int, len(accounts))
	idxMap := make(map[string]int)

	for idx, account := range accounts {
		uf[idx] = idx
		emails := account[1:]
		for _, email := range emails {
			if existingIdx, exists := idxMap[email]; exists {
				unionIndex(uf, existingIdx, idx)
				continue
			}
			idxMap[email] = idx
		}
	}

	resultMap := make(map[int][]string)
	for email, idx := range idxMap {
		root := findIndex(uf, idx)
		resultMap[root] = append(resultMap[root], email)
	}

	result := [][]string{}
	for idx, emails := range resultMap {
		sort.Strings(emails)
		result = append(result, append([]string{accounts[idx][0]}, emails...))
	}

	return result
}

func unionIndex(uf []int, x, y int) {
	rootX := findIndex(uf, x)
	rootY := findIndex(uf, y)
	if rootX != rootY {
		uf[rootX] = rootY
	}
}

func findIndex(uf []int, x int) int {
	if uf[x] != x {
		uf[x] = findIndex(uf, uf[x])
	}
	return uf[x]
}

/*
KP.
Use map to simulate forest.
*/
func accountsMergeV2(accounts [][]string) [][]string {

	result := [][]string{}
	parent := make(map[string]string)
	nameMap := make(map[string]string)

	for _, account := range accounts {
		name := account[0]
		emails := account[1:]

		for _, email := range emails {
			nameMap[email] = name
			if _, exists := parent[email]; !exists {
				parent[email] = email
			}
			union(parent, emails[0], email)
		}
	}

	groups := make(map[string][]string)
	for email := range parent {
		root := find(parent, email)
		groups[root] = append(groups[root], email)
	}
	for root, group := range groups {
		name := nameMap[root]
		sort.Strings(group)
		result = append(result, append([]string{name}, group...))
	}
	return result
}

func union(parent map[string]string, x, y string) {
	rootX := find(parent, x)
	rootY := find(parent, y)

	if rootX != rootY {
		parent[rootX] = rootY
	}
}

func find(parent map[string]string, x string) string {
	if parent[x] != x {
		parent[x] = find(parent, parent[x])
	}
	return parent[x]
}
