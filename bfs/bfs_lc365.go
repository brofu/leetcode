package bfs

import "fmt"

func canMeasureWater(x int, y int, target int) bool {

	getNeighbors := func(cx, cy int) [][]int {
		result := [][]int{
			{x, cy},            // full x
			{0, cy},            // empty x
			{0, (cx + cy) % y}, // from x to y
			{cx, y},            // full y
			{cx, 0},            // empty y
			{(cx + cy) % x, 0}, // from y to x
		}
		return result
	}

	visited := make(map[string]struct{})
	q := [][]int{{0, 0}}
	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {
			current := q[idx]
			if current[0]+current[1] == target {
				return true
			}
			neighbors := getNeighbors(current[0], current[1])
			for _, neighbor := range neighbors {
				if _, exist := visited[fmt.Sprintf("%d-%d", neighbor[0], neighbor[1])]; exist {
					continue
				}
				visited[fmt.Sprintf("%d-%d", neighbor[0], neighbor[1])] = struct{}{}
				q = append(q, neighbor)
			}
		}
		q = q[l:]
	}

	return false
}
