package interviews

import (
	"math"
	"strings"
)

/*
Agoda. Staff Software Engineer


Input:
	NY : LD
	NY : SG
	LD : BK
	SG : TK
	TK : BK

Output:
	NY : LD : BK



KP:
	1. Shortest path
	2. How to record the path? path[next] = current. Record backwards.

*/

func findPath(input []string, source, destination string) string {

	graph := make(map[string][]string)
	for _, p := range input {
		cities := strings.Split(p, ":")
		sour := strings.Trim(cities[0], " ")
		dest := strings.Trim(cities[1], " ")
		graph[sour] = append(graph[sour], dest)
	}

	path := make(map[string]string)
	visited := make(map[string]struct{})

	q := make([]string, 1)
	q[0] = source
	visited[source] = struct{}{}

breakFor:
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			current := q[i]
			if current == destination {
				break breakFor
			}
			for _, next := range graph[current] {
				if _, exist := visited[next]; exist {
					continue
				}
				visited[next] = struct{}{}
				path[next] = current
				q = append(q, next)
			}
		}
		q = q[l:]
	}

	result := []string{destination}
	for {
		pre, exist := path[destination]
		if exist {
			result = append(result, pre)
			destination = pre
		} else {
			break
		}
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return strings.Join(result, " : ")
}

/*

Output:
	return ALL the possible shorted path
*/
func findPathVar(input []string, source, destination string) []string {
	graph := make(map[string][]string)
	for _, p := range input {
		cities := strings.Split(p, ":")
		sour := strings.Trim(cities[0], " ")
		dest := strings.Trim(cities[1], " ")
		graph[sour] = append(graph[sour], dest)
	}

	type state struct {
		current       string
		distFromStart int
	}

	path := make(map[string][]string)
	distTo := make(map[string]int)

	q := make([]state, 1)
	q[0] = state{
		current:       source,
		distFromStart: 0,
	}
	for k, v := range graph {
		distTo[k] = math.MaxInt32
		for _, c := range v {
			distTo[c] = math.MaxInt32
		}
	}
	distTo[source] = 0

	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			current := q[i]
			if distTo[current.current] < current.distFromStart {
				continue
			}
			for _, next := range graph[current.current] {
				nextDist := current.distFromStart + 1
				if distTo[next] < nextDist {
					continue
				}
				if distTo[next] == nextDist {
					path[next] = append(path[next], current.current)
					continue
				}
				path[next] = []string{current.current}
				distTo[next] = nextDist
				q = append(q, state{
					current:       next,
					distFromStart: nextDist,
				})
			}
		}
		q = q[l:]
	}

	var buildResult func(string) [][]string

	buildResult = func(d string) [][]string {
		prevs, exist := path[d]
		if !exist { // base case
			if d == source {
				return [][]string{{d}} //todo
			}
			return [][]string{}
		}
		result := make([][]string, 0)
		for _, pre := range prevs {
			temp := buildResult(pre)
			for _, t := range temp {
				result = append(result, append([]string{d}, t...))
			}
		}
		return result
	}

	tempResults := buildResult(destination)
	for idx := range tempResults {
		result := tempResults[idx]
		for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
			result[i], result[j] = result[j], result[i]
		}
	}
	results := make([]string, len(tempResults))
	for idx, res := range tempResults {
		results[idx] = strings.Join(res, " : ")
	}

	return results
}

func findPathVarDFS(input []string, source, destination string) []string {
	graph := make(map[string][]string)
	for _, p := range input {
		cities := strings.Split(p, ":")
		sour := strings.TrimSpace(cities[0])
		dest := strings.TrimSpace(cities[1])
		graph[sour] = append(graph[sour], dest)
	}

	m := math.MaxInt32
	tempResults := make([][]string, 0)
	onPath := make(map[string]struct{})

	var dfs func(string, []string)
	dfs = func(source string, path []string) {

		// pruning
		if len(path) > m { // KP: there is already minimum result, can use for pruning
			return
		}
		// base case
		if source == destination {
			if len(path) < m {
				m = len(path)
				tempResults = [][]string{append([]string(nil), path...)}
				return
			}
			if len(path) == m {
				tempResults = append(tempResults, append([]string(nil), path...))
				return
			}
		}

		//KP: onPath v.s. visited? Check path or check nodes?
		if _, exist := onPath[source]; exist {
			return
		}

		onPath[source] = struct{}{}

		for _, next := range graph[source] {
			dfs(next, append(path, next))
		}

		delete(onPath, source)
	}
	dfs(source, []string{source})

	results := make([]string, len(tempResults))

	for idx := range tempResults {
		results[idx] = strings.Join(tempResults[idx], " : ")
	}
	return results
}
