package bfs

import (
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {

	result := [][]string{}
	emailMap := make(map[string][]int)
	for idx, emails := range accounts {
		for _, email := range emails[1:] {
			emailMap[email] = append(emailMap[email], idx)
		}
	}

	visitedMap := make(map[string]struct{})

	for email := range emailMap {

		if _, exists := visitedMap[email]; exists {
			continue
		}

		emailList := []string{}
		var accountName string

		q := []string{email}
		for len(q) > 0 {
			l := len(q)
			for idx := 0; idx < l; idx++ {
				current := q[idx]
				for _, accountIdx := range emailMap[current] {
					account := accounts[accountIdx]
					accountName = account[0]
					for i := 1; i < len(account); i++ {
						if _, exists := visitedMap[account[i]]; exists {
							continue
						}
						visitedMap[account[i]] = struct{}{}
						emailList = append(emailList, account[i])
						q = append(q, account[i])
					}
				}
			}
			q = q[l:]
		}
		sort.Strings(emailList)
		result = append(result, append([]string{accountName}, emailList...))
	}
	return result
}

/*
KP

TC
1. Build the graph. O(E). E is ALL the emails across all the accounts
2. Build the name map. O(E)
3. Traverse the graph. O(E+N). N is the unique email numbers. (E as edges and N as nodes)
4. Sort the result, O(NlgN)
5. Over all is O(E+N+NlogN), O(E+NlgN). E>=N always

SC
1. Graph O(E)
2. Name map O(N)
3. list O(N)
4. Queue size O(N)
5. Visited O(N)
6. Result O(N)
7. Overall O(E+4N) = O(E)
*/

func accountsMergeVersion2(accounts [][]string) [][]string {

	nameMap := make(map[string]string)
	visited := make(map[string]struct{})
	result := [][]string{}

	// construct the graph
	graph := make(map[string][]string)
	for _, list := range accounts {

		accountName := list[0]
		emails := list[1:]

		for _, email := range emails {
			nameMap[email] = accountName
		}
		graph[emails[0]] = append(graph[emails[0]], emails[1:]...)
		for _, email := range emails[1:] {
			graph[email] = append(graph[email], emails[0])
		}
	}

	for email := range graph {
		if _, exists := visited[email]; exists {
			continue
		}

		list := []string{email}
		visited[email] = struct{}{}

		q := []string{email}
		for len(q) > 0 {
			l := len(q)
			for idx := 0; idx < l; idx++ {
				current := q[idx]
				neighbors := graph[current]
				for _, neighbor := range neighbors {
					if _, exists := visited[neighbor]; exists {
						continue
					}
					list = append(list, neighbor)
					visited[neighbor] = struct{}{}
					q = append(q, neighbor)
				}
			}
			q = q[l:]
		}
		sort.Strings(list)
		result = append(result, append([]string{nameMap[email]}, list...))
	}

	return result
}
