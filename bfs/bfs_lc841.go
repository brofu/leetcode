package bfs

/*
KP

* TC
1. O(M*N). M is the number of rooms, and N is the average number of keys in each room

* SC
1. O(M). Queue size.
2. O(M). Visited array
3. Over all O(M)
*/
func canVisitAllRooms(rooms [][]int) bool {

	if len(rooms[0]) == 0 {
		return false
	}

	visited := make([]bool, len(rooms))
	visited[0] = true

	q := [][]int{rooms[0]}

	for len(q) > 0 {
		l := len(q)
		for idx := 0; idx < l; idx++ {
			current := q[idx]
			for _, room := range current {
				if visited[room] {
					continue
				}
				visited[room] = true
				keys := rooms[room]
				if len(keys) > 0 {
					q = append(q, keys)
				}
			}
		}
		q = q[l:]
	}

	for _, flag := range visited {
		if !flag {
			return false
		}
	}

	return true
}
